// Logging system designed for backend web application.
package wlog

import "github.com/axli-personal/wlog/mix"

// Global implementation.
var Mix = mix.NewSimpleLevelLogger(mix.Devp)
