package log

import (
	"bytes"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
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

	LevelEnabled
}

type LevelEnabled interface {
	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool
}

type stdLog struct {
	logrus *logrus.Logger
}

// NewStdLog creates a standard logger
func NewStdLog(level logrus.Level) Log {
	std := stdLog{}
	log := logrus.New()
	// 以Stdout为输出，代替默认的stderr
	log.SetOutput(os.Stdout)
	log.SetFormatter(&std)
	// 设置日志等级,环境变量配置优先
	parseLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err == nil {
		log.SetLevel(parseLevel)
	} else {
		log.SetLevel(level)
	}
	std.logrus = log
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
	level := entry.Level.String()
	if level == "warning" {
		level = "warn"
	}
	newLog := fmt.Sprintf("[%s] [ %s ] %s\n", timestamp, level, entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}

func (l *stdLog) Trace(args ...interface{}) {
	l.logrus.Trace(args...)
}

func (l *stdLog) TraceF(format string, args ...interface{}) {
	l.logrus.Tracef(format, args...)
}

func (l *stdLog) Debug(args ...interface{}) {
	l.logrus.Debug(args...)
}

func (l *stdLog) DebugF(format string, args ...interface{}) {
	l.logrus.Debugf(format, args...)
}

func (l *stdLog) Info(args ...interface{}) {
	l.logrus.Info(args...)
}

func (l *stdLog) InfoF(format string, args ...interface{}) {
	l.logrus.Infof(format, args...)
}

func (l *stdLog) Warn(args ...interface{}) {
	l.logrus.Warn(args...)
}

func (l *stdLog) WarnF(format string, args ...interface{}) {
	l.logrus.Warnf(format, args...)
}

func (l *stdLog) Error(args ...interface{}) {
	l.logrus.Error(args...)
}

func (l *stdLog) ErrorF(format string, args ...interface{}) {
	l.logrus.Errorf(format, args...)
}

func (l *stdLog) Fatal(args ...interface{}) {
	l.logrus.Fatal(args...)
}

func (l *stdLog) FatalF(format string, args ...interface{}) {
	l.logrus.Fatalf(format, args...)
}

func (l *stdLog) IsTraceEnabled() bool {
	return l.logrus.IsLevelEnabled(l.logrus.Level)
}

func (l *stdLog) IsDebugEnabled() bool {
	return l.logrus.IsLevelEnabled(l.logrus.Level)
}

func (l *stdLog) IsInfoEnabled() bool {
	return l.logrus.IsLevelEnabled(l.logrus.Level)
}

func (l *stdLog) IsWarnEnabled() bool {
	return l.logrus.IsLevelEnabled(l.logrus.Level)
}

func (l *stdLog) IsErrorEnabled() bool {
	return l.logrus.IsLevelEnabled(l.logrus.Level)
}

func (l *stdLog) IsFatalEnabled() bool {
	return l.logrus.IsLevelEnabled(l.logrus.Level)
}
