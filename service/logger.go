package service

import (
	"fmt"
	"github.com/axli-personal/wlog/base"
	"github.com/axli-personal/wlog/code"
	"io"
	"os"
)

// Logging in specific service.
type Logger interface {
	Print(url string, args ...interface{})
	Printf(url string, format string, args ...interface{})

	Fatal(url string, args ...interface{})
	Fatalf(url string, format string, args ...interface{})
}

type urlLogger struct {
	base.Logger
}

func NewLogger(out io.Writer, flag int) Logger {
	return &urlLogger{Logger: base.NewLogger(out, "SERVICE ", flag)}
}

func (service *urlLogger) Print(url string, args ...interface{}) {
	service.Logger.Output(1, url+"|"+fmt.Sprint(args...))
}

func (service *urlLogger) Printf(url string, format string, args ...interface{}) {
	service.Logger.Output(1, url+"|"+fmt.Sprintf(format, args...))
}

func (service *urlLogger) Fatal(url string, args ...interface{}) {
	service.Logger.Output(1, url+"|"+fmt.Sprint(args...))
	os.Exit(code.ServiceLoggerExit)
}

func (service *urlLogger) Fatalf(url string, format string, args ...interface{}) {
	service.Logger.Output(1, url+"|"+fmt.Sprintf(format, args...))
	os.Exit(code.ServiceLoggerExit)
}
