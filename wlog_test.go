package wlog

import (
	"github.com/axli-personal/wlog/mix"
	"testing"
)

func TestSimpleLevelLogger(t *testing.T) {
	Mix.Status(mix.Devp).Print("status information")
	Mix.Service(mix.Devp).Print("/api", "service information")
	Mix.Database(mix.Devp).Print("query", "table", "database information")
}
