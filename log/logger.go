package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// Create a new instance of the logger. You can have any number of instances.
var logger = logrus.New()

type MyFormatter struct {
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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

func init() {
	// 以JSON格式为输出，代替默认的ASCII格式
	//log.Formatter = &logrus.JSONFormatter{PrettyPrint: false}
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:            true, // 显示日志颜色
		DisableLevelTruncation: false,
		//QuoteEmptyFields: false,                 //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	}
	logger.Formatter = &MyFormatter{}
	// 以Stdout为输出，代替默认的stderr
	logger.SetOutput(os.Stdout)
	// 设置日志等级
	logger.SetLevel(logrus.TraceLevel)
}

func Trace(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}
