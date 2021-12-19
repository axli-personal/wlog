package simple

import (
	"fmt"
	"github.com/axli-personal/wlog/abstract"
	"github.com/axli-personal/wlog/code"
	"io"
	"log"
	"os"
)

type databaseLogger struct {
	logger abstract.GeneralLogger
}

func NewSimpleDatabaseLogger(out io.Writer, flag int) abstract.DatabaseLogger {
	database := new(databaseLogger)
	database.logger = log.New(out, "DATABASE ", flag)
	return database
}

func (database *databaseLogger) Print(operation string, table string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprint(args...))
}

func (database *databaseLogger) Printf(operation string, table string, format string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprintf(format, args...))
}

func (database *databaseLogger) Fatal(operation string, table string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprint(args...))
	os.Exit(code.DatabaseLoggerExit)
}

func (database *databaseLogger) Fatalf(operation string, table string, format string, args ...interface{}) {
	database.logger.Print(operation + "|" + table + "|" + fmt.Sprintf(format, args...))
	os.Exit(code.DatabaseLoggerExit)
}
