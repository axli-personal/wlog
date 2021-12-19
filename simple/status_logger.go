package simple

import (
	"fmt"
	"github.com/axli-personal/wlog/abstract"
	"github.com/axli-personal/wlog/code"
	"io"
	"log"
	"os"
)

type statusLogger struct {
	logger abstract.GeneralLogger
}

func NewSimpleStatusLogger(out io.Writer, flag int) abstract.StatusLogger {
	status := new(statusLogger)
	status.logger = log.New(out, "STATUS ", flag)
	return status
}

func (status *statusLogger) Print(args ...interface{}) {
	status.logger.Print(fmt.Sprint(args...))
}

func (status *statusLogger) Printf(format string, args ...interface{}) {
	status.logger.Print(fmt.Sprintf(format, args...))
}

func (status *statusLogger) Fatal(args ...interface{}) {
	status.logger.Print(fmt.Sprint(args...))
	os.Exit(code.StatusLoggerExit)
}

func (status *statusLogger) Fatalf(format string, args ...interface{}) {
	status.logger.Print(fmt.Sprintf(format, args...))
	os.Exit(code.StatusLoggerExit)
}
