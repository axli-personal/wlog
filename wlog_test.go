package wlog_test

import (
	"net/http"

	"github.com/axli-personal/wlog"
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	// Make logger.
	var logger wlog.Logger

	logger = wlog.NewLogger(os.Stdout, "||")
	logger = wlog.WithExtract(logger, "API")
	logger = wlog.WithExtract(logger, "Component")
	logger = wlog.WithExtract(logger, "Instance")
	logger = wlog.WithFlag(logger, wlog.Time|wlog.File)
	logger = wlog.WithMaxLevel(logger, wlog.Info, true)

	// Start logging.
	logger.OutputHeaders()

	request, err := http.NewRequest("GET", "10.10.10.10", nil)

	logger.Log(wlog.Pairs{
		{"Level", wlog.Debug},
		{"RequestContent", request},
		{"RequestError", err},
	})

	logger.Log(wlog.Pairs{
		{"Level", wlog.Info},
		{"Instance", 50},
		{"API", "/service/login"},
		{"IP", "100.100.100.100"},
	})

	logger.Log(wlog.Pairs{
		{"Level", wlog.Warn},
		{"Instance", 100},
		{"API", "/service/login"},
		{"Component", "Account"},
		{"Problem", "database fail to response"},
	})

	logger.Log(wlog.Pairs{
		{"Level", "Fatal"},
		{"Instance", 200},
		{"Problem", "memory not enough"},
		{"Memory", "8GB"},
	})
}
