package database

import (
	"fmt"
	"github.com/axli-personal/wlog/base"
	"github.com/axli-personal/wlog/code"
	"io"
	"os"
)

// Logger record details of database operation.
type Logger interface {
	Print(operation string, table string, args ...interface{})
	Printf(operation string, table string, format string, args ...interface{})

	Fatal(operation string, table string, args ...interface{})
	Fatalf(operation string, table string, format string, args ...interface{})
}

func NewLogger(out io.Writer, flag int) Logger {
	return &tableLogger{Logger: base.NewLogger(out, "DATABASE ", flag)}
}

type tableLogger struct {
	base.Logger
}

func (database *tableLogger) Print(operation string, table string, args ...interface{}) {
	database.Logger.Output(1, operation+"|"+table+"|"+fmt.Sprint(args...))
}

func (database *tableLogger) Printf(operation string, table string, format string, args ...interface{}) {
	database.Logger.Output(1, operation+"|"+table+"|"+fmt.Sprintf(format, args...))
}

func (database *tableLogger) Fatal(operation string, table string, args ...interface{}) {
	database.Logger.Output(1, operation+"|"+table+"|"+fmt.Sprint(args...))
	os.Exit(code.DatabaseLoggerExit)
}

func (database *tableLogger) Fatalf(operation string, table string, format string, args ...interface{}) {
	database.Logger.Output(1, operation+"|"+table+"|"+fmt.Sprintf(format, args...))
	os.Exit(code.DatabaseLoggerExit)
}
