// Logging system designed for backend web application.
package wlog

import (
	"log"
	"os"
)

const (
	Production = iota
	Development
)

var Mix = NewSimpleMixLevelLogger(Development)

type MixLevelLogger interface {
	ChangeLevel(level int)
	GetStatusLogger(level int) StatusLogger
	GetServiceLogger(level int) ServiceLogger
	GetDatabaseLogger(level int) DatabaseLogger
}

type SimpleMixLogger struct {
	production bool
	status     StatusLogger
	service    ServiceLogger
	database   DatabaseLogger
}

func NewSimpleMixLevelLogger(level int) MixLevelLogger {
	logger := new(SimpleMixLogger)
	logger.ChangeLevel(level)
	logger.status = NewSimpleStatusLogger(os.Stdout, log.Ldate|log.Ltime)
	logger.service = NewSimpleServiceLogger(os.Stdout, log.Ldate|log.Ltime)
	logger.database = NewSimpleDatabaseLogger(os.Stdout, log.Ldate|log.Ltime)
	return logger
}

func (mix *SimpleMixLogger) ChangeLevel(level int) {
	switch level {
	case Production:
		mix.production = true
	default:
		mix.production = false
	}
}

func (mix *SimpleMixLogger) GetStatusLogger(level int) StatusLogger {
	if mix.production && level != Production {
		return nil
	}
	return mix.status
}

func (mix *SimpleMixLogger) GetServiceLogger(level int) ServiceLogger {
	if mix.production && level != Production {
		return nil
	}
	return mix.service
}

func (mix *SimpleMixLogger) GetDatabaseLogger(level int) DatabaseLogger {
	if mix.production && level != Production {
		return nil
	}
	return mix.database
}
