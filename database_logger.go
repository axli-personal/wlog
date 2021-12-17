package wlog

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logging when perform database operation.
type DatabaseLogger interface {
	Print(operation string, table string, args ...interface{})
	Printf(operation string, table string, format string, args ...interface{})

	Fatal(operation string, table string, args ...interface{})
	Fatalf(operation string, table string, format string, args ...interface{})
}

type SimpleDatabaseLogger struct {
	logger GeneralLogger
}

func NewSimpleDatabaseLogger(out io.Writer, flag int) DatabaseLogger {
	database := new(SimpleDatabaseLogger)
	database.logger = log.New(out, "DATABASE ", flag)
	return database
}

func (database *SimpleDatabaseLogger) Print(operation string, table string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprint(args...))
}

func (database *SimpleDatabaseLogger) Printf(operation string, table string, format string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprintf(format, args...))
}

func (database *SimpleDatabaseLogger) Fatal(operation string, table string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprint(args...))
	os.Exit(DBLoggerExitCode)
}

func (database *SimpleDatabaseLogger) Fatalf(operation string, table string, format string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprintf(format, args...))
	os.Exit(DBLoggerExitCode)
}
