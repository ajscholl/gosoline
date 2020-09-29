package db_repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
	"strings"
)

type columnMetadata struct {
	exists     bool
	name       string
	nameQuoted string
	definition string
}

type tableMetadata struct {
	exists          bool
	tableName       string
	tableNameQuoted string
	columns         []columnMetadata
	primaryKeys     []columnMetadata
	unknownColumns  []columnMetadata
}

type tableMetadataBuilder struct {
	scope     *gorm.Scope
	tableName string
	fields    []*gorm.StructField
}

func (m *tableMetadataBuilder) build() *tableMetadata {
	metadata := &tableMetadata{}
	metadata.exists = m.scope.Dialect().HasTable(m.tableName)
	metadata.tableName = m.tableName
	metadata.tableNameQuoted = m.scope.Quote(m.tableName)
	metadata.columns = m.buildColumns()
	metadata.primaryKeys = m.buildPrimaryKeys()
	metadata.unknownColumns = m.buildUnknownColumns(metadata.columns, metadata.primaryKeys)
	return metadata
}

func (m *tableMetadataBuilder) buildColumns() []columnMetadata {
	var columns []columnMetadata
	for _, field := range m.fields {
		if field.IsNormal {
			columns = append(columns, m.buildColumn(field))
		}
	}
	return columns
}

func (m *tableMetadataBuilder) buildPrimaryKeys() []columnMetadata {
	var columns []columnMetadata
	for _, field := range m.fields {
		if field.IsPrimaryKey {
			columns = append(columns, m.buildColumn(field))
		}
	}
	return columns
}

type DescribeTableRow struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default *string
	Extra   string
}

func (m *tableMetadataBuilder) buildUnknownColumns(columns []columnMetadata, keys []columnMetadata) []columnMetadata {
	rows, err := m.scope.NewDB().Raw(fmt.Sprintf("DESCRIBE %s", m.scope.Quote(m.tableName))).Rows()

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var unknownColumns []columnMetadata
	for rows.Next() {
		row := DescribeTableRow{}
		err := rows.Scan(&row.Field, &row.Type, &row.Null, &row.Key, &row.Default, &row.Extra)

		if err != nil {
			panic(err)
		}

		fieldName := row.Field
		known := false

		for _, column := range columns {
			if column.name == fieldName {
				known = true
			}
		}
		for _, key := range keys {
			if key.name == fieldName {
				known = true
			}
		}

		if !known {
			definition := row.Type

			if row.Null == "YES" {
				definition += " NULL"
			} else {
				definition += " NOT NULL"
			}

			unknownColumns = append(unknownColumns, columnMetadata{
				exists:     true,
				name:       fieldName,
				nameQuoted: m.scope.Quote(fieldName),
				definition: definition,
			})
		}
	}
	return unknownColumns
}

func (m *tableMetadataBuilder) buildColumn(field *gorm.StructField) columnMetadata {
	return columnMetadata{
		name:       field.DBName,
		nameQuoted: m.scope.Quote(field.DBName),
		definition: m.scope.Quote(field.DBName) + " " + m.dataTypeOfField(field),
		exists:     m.scope.Dialect().HasColumn(m.tableName, field.DBName),
	}
}

func (m *tableMetadataBuilder) dataTypeOfField(field *gorm.StructField) string {
	tag := m.scope.Dialect().DataTypeOf(field)

	tag = strings.Replace(tag, "AUTO_INCREMENT", "", -1)
	tag = strings.Replace(tag, "UNIQUE", "", -1)

	return tag
}

func newTableMetadata(scope *gorm.Scope, tableName string, fields []*gorm.StructField) *tableMetadata {
	builder := tableMetadataBuilder{
		tableName: tableName,
		scope:     scope,
		fields:    fields,
	}
	return builder.build()
}

func (m *tableMetadata) columnNamesQuoted() []string {
	return m.namesQuoted(m.columns)
}

func (m *tableMetadata) primaryKeyNamesQuoted() []string {
	return m.namesQuoted(m.primaryKeys)
}

func (m *tableMetadata) columnDefinitions() []string {
	return m.definitions(m.columns)
}

func (m *tableMetadata) namesQuoted(items []columnMetadata) []string {
	return funk.Map(items, func(item columnMetadata) string {
		return item.nameQuoted
	}).([]string)
}

func (m *tableMetadata) definitions(items []columnMetadata) []string {
	return funk.Map(items, func(item columnMetadata) string {
		return item.definition
	}).([]string)
}

func (m *tableMetadata) columnNamesQuotedExcludingValue(excluded ...string) []string {
	return m.namesQuoted(funk.Filter(m.columns, func(item columnMetadata) bool {
		return !funk.ContainsString(excluded, item.name)
	}).([]columnMetadata))
}
