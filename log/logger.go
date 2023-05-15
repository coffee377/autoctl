package log

import (
	"github.com/coffee377/autoctl/lib"
	"github.com/sirupsen/logrus"
)

var logger = lib.NewStdLog(logrus.TraceLevel)

func Trace(format string, v ...interface{}) {
	logger.TraceF(format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.DebugF(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.InfoF(format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.WarnF(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.ErrorF(format, v...)
}

func Fatal(format string, v ...interface{}) {
	logger.FatalF(format, v...)
}
