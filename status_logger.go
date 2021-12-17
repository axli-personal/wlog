package wlog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type StatusLogger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type SimpleStatusLogger struct {
	logger GeneralLogger
}

func NewSimpleStatusLogger(out io.Writer, flag int) StatusLogger {
	status := new(SimpleStatusLogger)
	status.logger = log.New(out, "STATUS ", flag)
	return status
}

func (status *SimpleStatusLogger) Print(args ...interface{}) {
	status.logger.Print(fmt.Sprint(args...))
}

func (status *SimpleStatusLogger) Printf(format string, args ...interface{}) {
	status.logger.Print(fmt.Sprintf(format, args...))
}

func (status *SimpleStatusLogger) Fatal(args ...interface{}) {
	status.logger.Print(fmt.Sprint(args...))
	os.Exit(StatusLoggerExitCode)
}

func (status *SimpleStatusLogger) Fatalf(format string, args ...interface{}) {
	status.logger.Print(fmt.Sprintf(format, args...))
	os.Exit(StatusLoggerExitCode)
}
