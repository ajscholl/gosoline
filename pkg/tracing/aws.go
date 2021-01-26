package tracing

import (
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/mon"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
	"strings"
)

func AWS(config cfg.Config, c *client.Client) {
	enabled := config.GetBool("tracing_enabled")

	if !enabled {
		return
	}

	xray.AWS(c)
}

type xrayLogger struct {
	logger mon.Logger
}

func newXrayLogger(logger mon.Logger) *xrayLogger {
	return &xrayLogger{
		logger: logger.WithChannel("tracing"),
	}
}

func (x xrayLogger) Log(level xraylog.LogLevel, msg fmt.Stringer) {
	msgStr := msg.String()

	switch level {
	case xraylog.LogLevelDebug:
		if strings.HasPrefix(msgStr, "{") {
			// debug statement at github.com/aws/aws-xray-sdk-go@v1.1.0/xray/default_emitter.go:67 logging the packet itself
			x.logger.Infof("about to send packet '%s'", msgStr)
		} else {
			x.logger.Debug(msgStr)
		}
	case xraylog.LogLevelInfo:
		x.logger.Info(msgStr)
	case xraylog.LogLevelWarn:
		x.logger.Warn(msgStr)
	case xraylog.LogLevelError:
		x.logger.Error(fmt.Errorf(msgStr), msgStr)
	}
}
