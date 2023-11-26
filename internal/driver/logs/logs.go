package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/elvenworks/prohermes"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
	"go.elastic.co/apm/module/apmlogrus/v2"
)

func Init() {
	InitLogrus()
}

func InitLogrus() {
	formater := &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	l := logrus.WithFields(logrus.Fields{})
	l.Logger.AddHook(&apmlogrus.Hook{})

	promermes := prohermes.MustNewPrometheusHook()

	l.Logger.SetReportCaller(true)
	l.Logger.SetFormatter(formater)
	l.Logger.AddHook(promermes)

	l.Logger.SetLevel(GetLoggerLevel())
}

func GetLoggerLevel() logrus.Level {
	loggerLevel := os.Getenv("LOGGER_LEVEL")

	switch loggerLevel {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.DebugLevel
	}
}
