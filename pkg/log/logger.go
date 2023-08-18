package log

import (
	"github.com/sirupsen/logrus"
)

var logger = NewStdLog(logrus.InfoLevel)

func Trace(format string, v ...interface{}) {
	logger.TraceF(format, v...)
}

func IsTraceEnabled() bool {
	return logger.IsTraceEnabled()
}

func Debug(format string, v ...interface{}) {
	logger.DebugF(format, v...)
}

func IsDebugEnabled() bool {
	return logger.IsDebugEnabled()
}

func Info(format string, v ...interface{}) {
	logger.InfoF(format, v...)
}

func IsInfoEnabled() bool {
	return logger.IsInfoEnabled()
}

func Warn(format string, v ...interface{}) {
	logger.WarnF(format, v...)
}

func IsWarnEnabled() bool {
	return logger.IsWarnEnabled()
}

func Error(format string, v ...interface{}) {
	logger.ErrorF(format, v...)
}

func IsErrorEnabled() bool {
	return logger.IsErrorEnabled()
}

func Fatal(format string, v ...interface{}) {
	logger.FatalF(format, v...)
}

func IsFatalEnabled() bool {
	return logger.IsFatalEnabled()
}
