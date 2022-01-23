package base

import (
	"io"
	"log"
)

const testMessage = "logging system output test"

// Logger should output message in concurrent environment.
//
// Test should raise error when logger is in bad status.
type Logger interface {
	Output(callDepth int, msg string)
	Test(callDepth int) error
}

func NewLogger(out io.Writer, prefix string, flag int) Logger {
	return &stdLogger{Logger: log.New(out, prefix, flag)}
}

type stdLogger struct {
	*log.Logger
}

func (l *stdLogger) Output(callDepth int, msg string) {
	_ = l.Logger.Output(callDepth+2, msg)
}

func (l *stdLogger) Test(callDepth int) error {
	return l.Logger.Output(callDepth+2, testMessage)
}
