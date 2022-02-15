package wlog_test

import (
	"errors"

	"github.com/axli-personal/wlog"
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	var logger wlog.Logger

	logger = wlog.NewLogger(os.Stdout, "||")
	logger = wlog.WithMaxLevel(logger, wlog.Info, true)
	logger = wlog.WithExtract(logger, "Instance")
	logger = wlog.WithExtract(logger, "API")
	logger = wlog.WithExtract(logger, "Component")
	logger = wlog.WithFlag(logger, wlog.Time|wlog.File)

	logger.MakeHeaders(nil)

	err := errors.New("database fail to response")

	logger.Log(wlog.Options{
		{"level", wlog.Warn},
		{"Instance", "100"},
		{"API", "/service/login"},
		{"Component", "Account"},
		{"Problem", err},
	})

	logger.Log(wlog.Options{
		{"level", "Fatal"},
		{"Instance", "200"},
		{"Problem", "memory not enough"},
		{"Memory", "8GB"},
	})
}
