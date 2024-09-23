package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/hashicorp/go-multierror"
	"github.com/jmoiron/sqlx"
	"github.com/justtrackio/gosoline/pkg/appctx"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/exec"
	"github.com/justtrackio/gosoline/pkg/log"
)

const (
	FormatDateTime = "2006-01-02 15:04:05"
)

//go:generate mockery --name SqlResult
type SqlResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Sqler interface {
	ToSql() (string, []any, error)
}

func SqlFmt(format string, formatValues []any, args ...any) Sqler {
	return sqlFmt{
		format:       format,
		formatValues: formatValues,
		args:         args,
	}
}

type sqlFmt struct {
	format       string
	formatValues []any
	args         []any
}

func (s sqlFmt) ToSql() (qry string, args []any, err error) {
	qry = fmt.Sprintf(s.format, s.formatValues...)
	args = s.args

	return
}

type (
	ResultRow map[string]string
	Result    []ResultRow
)

//go:generate mockery --name Client
type Client interface {
	GetSingleScalarValue(ctx context.Context, query string, args ...any) (int, error)
	GetResult(ctx context.Context, query string, args ...any) (*Result, error)
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExec(ctx context.Context, query string, arg any) (sql.Result, error)
	ExecMultiInTx(ctx context.Context, sqlers ...Sqler) (results []sql.Result, err error)
	BindNamed(query string, arg any) (string, []any, error)
	Prepare(ctx context.Context, query string) (*sql.Stmt, error)
	Preparex(ctx context.Context, query string) (*sqlx.Stmt, error)
	PrepareNamed(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	Queryx(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
	NamedQuery(ctx context.Context, query string, arg any) (*sqlx.Rows, error)
	Select(ctx context.Context, dest any, query string, args ...any) error
	NamedSelect(ctx context.Context, dest any, query string, arg any) error
	Get(ctx context.Context, dest any, query string, args ...any) error
	WithTx(ctx context.Context, ops *sql.TxOptions, do func(ctx context.Context, tx *sql.Tx) error) error
}

type ClientSqlx struct {
	logger   log.Logger
	db       *sqlx.DB
	executor exec.Executor
}

type clientCtxKey string

func ProvideClient(ctx context.Context, config cfg.Config, logger log.Logger, name string) (Client, error) {
	var err error
	var settings *Settings

	if settings, err = readSettings(config, name); err != nil {
		return nil, err
	}

	return appctx.Provide(ctx, clientCtxKey(fmt.Sprint(settings)), func() (Client, error) {
		return NewClientWithSettings(ctx, config, logger, name, settings)
	})
}

func NewClient(ctx context.Context, config cfg.Config, logger log.Logger, name string) (Client, error) {
	var err error
	var settings *Settings

	if settings, err = readSettings(config, name); err != nil {
		return nil, err
	}

	return NewClientWithSettings(ctx, config, logger, name, settings)
}

func NewClientWithSettings(ctx context.Context, config cfg.Config, logger log.Logger, name string, settings *Settings) (Client, error) {
	var err error
	var connection *sqlx.DB
	var executor exec.Executor = exec.NewDefaultExecutor()

	if connection, err = ProvideConnectionFromSettings(ctx, logger, settings); err != nil {
		return nil, fmt.Errorf("can not connect to sql database: %w", err)
	}

	if settings.Retry.Enabled {
		path := fmt.Sprintf("db.%s.retry", name)
		executor = NewExecutor(config, logger, name, path)
	}

	return NewClientWithInterfaces(logger, connection, executor), nil
}

func NewClientWithInterfaces(logger log.Logger, connection *sqlx.DB, executor exec.Executor) Client {
	return &ClientSqlx{
		logger:   logger,
		db:       connection,
		executor: executor,
	}
}

func (c *ClientSqlx) GetSingleScalarValue(ctx context.Context, query string, args ...any) (int, error) {
	var val sql.NullInt64
	err := c.Get(ctx, &val, query, args...)
	if err != nil {
		return 0, err
	}

	if !val.Valid {
		return 0, nil
	}

	return int(val.Int64), err
}

func (c *ClientSqlx) GetResult(ctx context.Context, query string, args ...any) (*Result, error) {
	out := make(Result, 0, 32)
	rows, err := c.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("can not get columns: %w", err)
	}

	types := make(map[string]string)

	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m := make(ResultRow)
		for i, colName := range cols {
			val := columnPointers[i].(*any)

			if _, ok := types[colName]; !ok {
				types[colName] = reflect.TypeOf(*val).String()
			}

			switch types[colName] {
			case "string":
				m[colName] = (*val).(string)
			case "[]uint8":
				m[colName] = string((*val).([]uint8))
			case "int":
				m[colName] = strconv.FormatInt(int64((*val).(int)), 10)
			case "int64":
				m[colName] = strconv.FormatInt((*val).(int64), 10)
			case "float64":
				m[colName] = strconv.FormatFloat((*val).(float64), 'f', -1, 64)
			default:
				errStr := fmt.Sprintf("could not convert mysql result into string map: %v -> %v is %v", colName, *val, reflect.TypeOf(*val))

				return nil, errors.New(errStr)
			}
		}

		out = append(out, m)
	}

	return &out, err
}

func (c *ClientSqlx) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	c.logger.Debug("> %s %q", query, args)

	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.ExecContext(ctx, query, args...)
	})
	if err != nil {
		return nil, err
	}

	return res.(sql.Result), err
}

func (c *ClientSqlx) NamedExec(ctx context.Context, query string, arg any) (sql.Result, error) {
	c.logger.Debug("> %s %q", query, arg)

	return c.db.NamedExecContext(ctx, query, arg)
}

func (c *ClientSqlx) ExecMultiInTx(ctx context.Context, sqlers ...Sqler) (results []sql.Result, err error) {
	var tx *sql.Tx
	var res sql.Result
	var queries []string
	var argss [][]any

	for i, sqler := range sqlers {
		var buildErr error
		var qry string
		var args []any

		if qry, args, buildErr = sqler.ToSql(); buildErr != nil {
			return nil, fmt.Errorf("can not build sql #%d: %w", i, err)
		}

		queries = append(queries, qry)
		argss = append(argss, args)
	}

	if tx, err = c.BeginTx(ctx, &sql.TxOptions{}); err != nil {
		return nil, fmt.Errorf("can not begin transaction: %w", err)
	}

	defer func() {
		if err == nil {
			return
		}

		if errRollback := tx.Rollback(); errRollback != nil {
			err = multierror.Append(err, fmt.Errorf("can not rollback tx: %w", errRollback)).ErrorOrNil()

			return
		}
	}()

	for i, qry := range queries {
		if res, err = c.Exec(ctx, qry, argss[i]...); err != nil {
			return nil, fmt.Errorf("can not exec qry %s: %w", qry, err)
		}

		results = append(results, res)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("can not commit transaction: %w", err)
	}

	return
}

func (c *ClientSqlx) BindNamed(query string, arg any) (qry string, args []any, err error) {
	return c.db.BindNamed(query, arg)
}

func (c *ClientSqlx) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.PrepareContext(ctx, query)
	})
	if err != nil {
		return nil, err
	}

	return res.(*sql.Stmt), nil
}

func (c *ClientSqlx) Preparex(ctx context.Context, query string) (*sqlx.Stmt, error) {
	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.PreparexContext(ctx, query)
	})
	if err != nil {
		return nil, err
	}

	return res.(*sqlx.Stmt), nil
}

func (c *ClientSqlx) PrepareNamed(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.PrepareNamedContext(ctx, query)
	})
	if err != nil {
		return nil, err
	}

	return res.(*sqlx.NamedStmt), nil
}

func (c *ClientSqlx) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	c.logger.Debug("> %s %q", query, args)

	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.QueryContext(ctx, query, args...)
	})
	if err != nil {
		return nil, err
	}

	return res.(*sql.Rows), nil
}

func (c *ClientSqlx) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	c.logger.Debug("> %s %q", query, args)

	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.QueryRowContext(ctx, query, args...), nil
	})
	if err != nil {
		return nil
	}

	return res.(*sql.Row)
}

func (c *ClientSqlx) NamedQuery(ctx context.Context, query string, arg any) (*sqlx.Rows, error) {
	c.logger.Debug("> %s %q", query, arg)

	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.NamedQueryContext(ctx, query, arg)
	})
	if err != nil {
		return nil, err
	}

	return res.(*sqlx.Rows), nil
}

func (c *ClientSqlx) Queryx(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	c.logger.Debug("> %s %q", query, args)

	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.QueryxContext(ctx, query, args...)
	})
	if err != nil {
		return nil, err
	}

	return res.(*sqlx.Rows), err
}

func (c *ClientSqlx) Select(ctx context.Context, dest any, query string, args ...any) error {
	c.logger.Debug("> %s %q", query, args)

	_, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return nil, c.db.SelectContext(ctx, dest, query, args...)
	})

	return err
}

func (c *ClientSqlx) NamedSelect(ctx context.Context, dest any, query string, arg any) error {
	c.logger.Debug("> %s %q", query, arg)

	_, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		stmt, err := c.db.PrepareNamedContext(ctx, query)
		if err != nil {
			return nil, err
		}
		defer func() {
			err = stmt.Close()
			if err != nil {
				c.logger.Error("can not close named statement: %w", err)
			}
		}()

		return nil, stmt.SelectContext(ctx, dest, arg)
	})

	return err
}

func (c *ClientSqlx) Get(ctx context.Context, dest any, query string, args ...any) error {
	c.logger.Debug("> %s %q", query, args)

	_, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return nil, c.db.GetContext(ctx, dest, query, args...)
	})

	return err
}

func (c *ClientSqlx) BeginTx(ctx context.Context, ops *sql.TxOptions) (*sql.Tx, error) {
	c.logger.Debug("start tx")

	res, err := c.executor.Execute(ctx, func(ctx context.Context) (any, error) {
		return c.db.BeginTx(ctx, ops)
	})
	if err != nil {
		return nil, err
	}

	return res.(*sql.Tx), err
}

func (c *ClientSqlx) WithTx(ctx context.Context, ops *sql.TxOptions, do func(ctx context.Context, tx *sql.Tx) error) (err error) {
	var tx *sql.Tx
	tx, err = c.BeginTx(ctx, ops)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				err = multierror.Append(err, fmt.Errorf("can not rollback tx: %w", errRollback))

				return
			}
			c.logger.WithContext(ctx).Debug("rollback successfully done")
		}
	}()

	err = do(ctx, tx)
	if err != nil {
		return fmt.Errorf("can not execute do function: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("can not commit tx: %w", err)
	}

	return nil
}
