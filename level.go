package wlog

import (
	"strconv"
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

// WithMaxLevel will wrap a logger and enable it to detect levels.
//
// A negative maxLevel will turn off the logger.
//
// Passing in nil logger will cause panic.
func WithMaxLevel(logger Logger, maxLevel int, format bool) Logger {
	if logger == nil {
		panic("can't create logger from nil")
	}

	IncDepth(logger)

	return &levelFilter{maxLevel: maxLevel, format: format, logger: logger}
}

type levelFilter struct {
	maxLevel int
	format   bool
	logger   Logger
}

func (filter *levelFilter) Log(pairs Pairs, columns ...string) {
	logLevel := Unset

	for i := 0; i < len(pairs); i++ {
		if strings.ToLower(pairs[i].Key) == "level" {
			switch level := pairs[i].Val.(type) {
			case string:
				logLevel = getIntLevel(level)
			case int:
				if level < Unset || level > Debug {
					logLevel = Unset
				} else {
					logLevel = level
				}
			}

			pairs[i].Key, pairs[i].Val = "", nil
			break
		}
	}

	if logLevel > filter.maxLevel {
		return
	}

	if filter.format {
		columns = append(columns, getStrLevel(logLevel))
	} else {
		columns = append(columns, strconv.Itoa(logLevel))
	}

	filter.logger.Log(pairs, columns...)

	if logLevel == Fatal {
		panic("log with fatal level")
	}
}

func (filter *levelFilter) OutputHeaders(headers ...string) {
	headers = append(headers, "Level")
	filter.logger.OutputHeaders(headers...)
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

func (filter *levelFilter) Next() Logger {
	return filter.logger
}
