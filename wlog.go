// Logging system designed for backend web application.
package wlog

import (
	"github.com/axli-personal/wlog/mix"
	"log"
	"os"
)

// Global implementation.
var Mix = mix.NewLevelLogger(os.Stdout, log.Ldate|log.Ltime|log.Lshortfile, mix.Devp)
