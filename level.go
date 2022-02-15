package wlog

import (
	"os"
	"strconv"
	"strings"
)

// These levels are supported in the system
// and most of them are picked from log4j.
//
// The logger with higher level will record more message.
const (
	Off = iota - 1
	Unset
	Fatal
	Error
	Warn
	Info
	Debug
)

// WithMaxLevel will wrapper a logger and enhance its ability.
//
// The level will not be preprocessed, so negative level will
// block every message and level greater than Debug will record
// every message. This is the default behavior of the system.
func WithMaxLevel(logger Logger, level int, format bool) Logger {
	if logger == nil {
		panic("can't create logger from nil")
	}

	IncDepth(logger)

	return &levelFilter{maxLevel: level, format: format, logger: logger}
}

type levelFilter struct {
	maxLevel int
	format   bool
	logger   Logger
}

func (filter *levelFilter) Log(options Options, columns ...string) {
	logLevel := Unset

	for i := 0; i < len(options); i++ {
		if strings.ToLower(options[i].Key) == "level" {
			switch level := options[i].Val.(type) {
			case string:
				logLevel = getIntLevel(level)
			case int:
				if level < Unset || level > Debug {
					logLevel = Unset
				} else {
					logLevel = level
				}
			}

			options[i].Key, options[i].Val = "", nil
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

	filter.logger.Log(options, columns...)

	if logLevel == Fatal {
		os.Exit(1)
	}
}

func (filter *levelFilter) MakeHeaders(headers []string) {
	headers = append(headers, "Level")
	filter.logger.MakeHeaders(headers)
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
