package log

import (
	"github.com/anqiansong/ketty/console"
	"github.com/anqiansong/ketty/text"
)

var c = console.NewConsole(console.WithTextOption(text.DisableBorder()))

func Debug(format string, v ...interface{}) {
	c.Debug(format, v)
}

func Info(format string, v ...interface{}) {
	c.Info(format, v)
}

func Warn(format string, v ...interface{}) {
	c.Warn(format, v)
}

func Error(err error) {
	c.Error(err)
}

func ErrorText(format string, v ...interface{}) {
	c.ErrorText(format, v)
}
