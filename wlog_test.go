package wlog

import "testing"

func TestSimpleMixLogger(t *testing.T) {
	if logger := Mix.GetStatusLogger(Development); logger != nil {
		logger.Print("status information")
	}
	if logger := Mix.GetServiceLogger(Development); logger != nil {
		logger.Print("/api", "service information")
	}
	if logger := Mix.GetDatabaseLogger(Development); logger != nil {
		logger.Print("query", "table", "database information")
	}
}
