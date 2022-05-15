package wlog_test

import (
	"github.com/axli-personal/wlog"
	"net/http"
	"os"
)

func Example() {
	// Make logger.
	var logger wlog.Logger

	logger = wlog.NewLogger(os.Stdout, "||")
	logger = wlog.WithExtract(logger, "API")
	logger = wlog.WithExtract(logger, "Component")
	logger = wlog.WithExtract(logger, "Instance")
	logger = wlog.WithFlag(logger, wlog.File)
	logger = wlog.WithMaxLevel(logger, wlog.Info)

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

	defer func() { recover() }()
	logger.Log(wlog.Pairs{
		{"Level", "Fatal"},
		{"Instance", 200},
		{"Problem", "memory not enough"},
		{"Memory", "8GB"},
	})

	// Output:
	// Level||File||Instance||Component||API||Details
	// Info||example_test.go:31||50||NULL||/service/login||IP=100.100.100.100
	// Warn||example_test.go:38||100||Account||/service/login||Problem=database fail to response
	// Fatal||example_test.go:47||200||NULL||NULL||Problem=memory not enough;Memory=8GB
}
