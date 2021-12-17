package wlog

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logging in specific service.
type ServiceLogger interface {
	Print(url string, args ...interface{})
	Printf(url string, format string, args ...interface{})

	Fatal(url string, args ...interface{})
	Fatalf(url string, format string, args ...interface{})
}

type SimpleServiceLogger struct {
	logger GeneralLogger
}

func NewSimpleServiceLogger(out io.Writer, flag int) ServiceLogger {
	service := new(SimpleServiceLogger)
	service.logger = log.New(out, "SERVICE ", flag)
	return service
}

func (service *SimpleServiceLogger) Print(url string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprint(args...))
}

func (service *SimpleServiceLogger) Printf(url string, format string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprintf(format, args...))
}

func (service *SimpleServiceLogger) Fatal(url string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprint(args...))
	os.Exit(ServiceLoggerExitCode)
}

func (service *SimpleServiceLogger) Fatalf(url string, format string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprintf(format, args...))
	os.Exit(ServiceLoggerExitCode)
}
