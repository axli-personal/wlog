package wlog

import (
	"strings"
)

// These are the levels available in the system, mostly from log4j.
const (
	Off = iota - 1
	Unset
	Fatal // Log with this level will call panic().
	Error
	Warn
	Info
	Debug
)

const DefaultMaxLevel = Info

// WithMaxLevel will wrap a logger and enable it to detect levels.
//
// A negative maxLevel will turn off the logger
// and Unset(0) will be converted to DefaultMaxLevel.
// You can use ParseLevel to parse level from string and integer.
//
// Passing in nil logger will cause panic.
func WithMaxLevel(logger Logger, maxLevel int) Logger {
	if logger == nil {
		panic("can't create logger from nil")
	}

	if maxLevel == Unset {
		maxLevel = DefaultMaxLevel
	}

	IncDepth(logger)

	return &levelFilter{maxLevel: maxLevel, logger: logger}
}

type levelFilter struct {
	maxLevel int
	logger   Logger
}

func (filter *levelFilter) Log(pairs Pairs, columns ...string) {
	logLevel := Unset

	for i := 0; i < len(pairs); i++ {
		if strings.ToLower(pairs[i].Key) == "level" {
			logLevel = ParseLevel(pairs[i].Val)
			if logLevel < Unset || logLevel > Debug {
				logLevel = Unset
			}

			pairs[i].Key, pairs[i].Val = "", nil
			break
		}
	}

	if logLevel > filter.maxLevel {
		return
	}

	columns = append(columns, getStrLevel(logLevel))

	filter.logger.Log(pairs, columns...)

	if logLevel == Fatal {
		panic("log with fatal level")
	}
}

func (filter *levelFilter) OutputHeaders(headers ...string) {
	headers = append(headers, "Level")
	filter.logger.OutputHeaders(headers...)
}

func (filter *levelFilter) Next() Logger {
	return filter.logger
}

// ParseLevel can parse level from string and integer.
func ParseLevel(arg interface{}) int {
	switch level := arg.(type) {
	case int:
		return level
	case string:
		return getIntLevel(level)
	default:
		return Unset
	}
}

func getIntLevel(level string) int {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	case "off":
		return Off
	default:
		return Unset
	}
}

func getStrLevel(level int) string {
	switch level {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "Unset"
	}
}
