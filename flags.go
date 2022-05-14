package wlog

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

const DefaultTimeFormat = "2006-01-02 15:04:05.000"

// These are the flags available in the system, mostly from standard log package.
const (
	Time = 1 << iota
	File
)

// WithFlag will wrap a logger and enable it to output useful messages.
//
// Passing in nil logger will cause panic.
func WithFlag(logger Logger, flag int) Logger {
	if logger == nil {
		panic("can't create logger from nil")
	}

	return &flagCollector{flag: flag, depth: 1, logger: logger}
}

type flagCollector struct {
	flag   int
	depth  int
	logger Logger
}

func (collector *flagCollector) Log(options Pairs, columns ...string) {
	if collector.flag&Time != 0 {
		columns = append(columns, time.Now().Format(DefaultTimeFormat))
	}

	if collector.flag&File != 0 {
		_, filePath, line, _ := runtime.Caller(collector.depth)
		_, fileName := path.Split(filePath)
		str := fmt.Sprintf("%s:%d", fileName, line)
		columns = append(columns, str)
	}

	collector.logger.Log(options, columns...)
}

func (collector *flagCollector) OutputHeaders(headers ...string) {
	if collector.flag&Time != 0 {
		headers = append(headers, "Time")
	}
	if collector.flag&File != 0 {
		headers = append(headers, "File")
	}
	collector.logger.OutputHeaders(headers...)
}

func IncDepth(logger Logger) {
	// This may be expensive.
	for logger != nil {
		if l, ok := logger.(*flagCollector); ok {
			l.depth++
		}
		logger = logger.Next()
	}
}

func (collector *flagCollector) Next() Logger {
	return collector.logger
}
