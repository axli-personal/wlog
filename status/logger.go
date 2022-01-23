package status

import (
	"fmt"
	"github.com/axli-personal/wlog/base"
	"github.com/axli-personal/wlog/code"
	"io"
	"os"
)

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type strLogger struct {
	base.Logger
}

func NewLogger(out io.Writer, flag int) Logger {
	return &strLogger{Logger: base.NewLogger(out, "STATUS ", flag)}
}

func (status *strLogger) Print(args ...interface{}) {
	status.Logger.Output(1, fmt.Sprint(args...))
}

func (status *strLogger) Printf(format string, args ...interface{}) {
	status.Logger.Output(1, fmt.Sprintf(format, args...))
}

func (status *strLogger) Fatal(args ...interface{}) {
	status.Logger.Output(1, fmt.Sprint(args...))
	os.Exit(code.StatusLoggerExit)
}

func (status *strLogger) Fatalf(format string, args ...interface{}) {
	status.Logger.Output(1, fmt.Sprintf(format, args...))
	os.Exit(code.StatusLoggerExit)
}
