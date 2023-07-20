package lib

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type Log interface {
	Trace(args ...interface{})
	TraceF(format string, args ...interface{})
	Debug(args ...interface{})
	DebugF(format string, args ...interface{})
	Info(args ...interface{})
	InfoF(format string, args ...interface{})
	Warn(args ...interface{})
	WarnF(format string, args ...interface{})
	Error(args ...interface{})
	ErrorF(format string, args ...interface{})
	Fatal(args ...interface{})
	FatalF(format string, args ...interface{})
}

type stdLog struct {
	log *logrus.Logger
}

// NewStdLog creates a standard logger
func NewStdLog(level logrus.Level) Log {
	std := stdLog{}
	log := logrus.New()
	// 以Stdout为输出，代替默认的stderr
	log.SetOutput(os.Stdout)
	// 设置日志等级,环境变量配置优先
	parseLevel, err := logrus.ParseLevel(os.Getenv("LOGGER_LEVEL"))
	if err == nil {
		log.SetLevel(parseLevel)
	} else {
		log.SetLevel(level)
	}
	std.log = log
	return &std
}

// Format 日志格式化
func (l *stdLog) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006/01/02 15:04:05")
	var newLog string
	newLog = fmt.Sprintf("[%s] [\t%s\t] %s\n", timestamp, entry.Level.String(), entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}

func (l *stdLog) Trace(args ...interface{}) {
	l.log.Trace(args...)
}

func (l *stdLog) TraceF(format string, args ...interface{}) {
	l.log.Tracef(format, args...)
}

func (l *stdLog) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *stdLog) DebugF(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l *stdLog) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *stdLog) InfoF(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *stdLog) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *stdLog) WarnF(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *stdLog) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *stdLog) ErrorF(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *stdLog) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l *stdLog) FatalF(format string, args ...interface{}) {
	l.log.Fatalf(format, args...)
}
