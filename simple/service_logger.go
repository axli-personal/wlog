package simple

import (
	"fmt"
	"github.com/axli-personal/wlog/abstract"
	"github.com/axli-personal/wlog/code"
	"io"
	"log"
	"os"
)

type serviceLogger struct {
	logger abstract.GeneralLogger
}

func NewSimpleServiceLogger(out io.Writer, flag int) abstract.ServiceLogger {
	service := new(serviceLogger)
	service.logger = log.New(out, "SERVICE ", flag)
	return service
}

func (service *serviceLogger) Print(url string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprint(args...))
}

func (service *serviceLogger) Printf(url string, format string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprintf(format, args...))
}

func (service *serviceLogger) Fatal(url string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprint(args...))
	os.Exit(code.ServiceLoggerExit)
}

func (service *serviceLogger) Fatalf(url string, format string, args ...interface{}) {
	service.logger.Print(url + "|" + fmt.Sprintf(format, args...))
	os.Exit(code.ServiceLoggerExit)
}
