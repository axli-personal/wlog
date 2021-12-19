package mix

import (
	"github.com/axli-personal/wlog/abstract"
	"github.com/axli-personal/wlog/null"
	"github.com/axli-personal/wlog/simple"
	"log"
	"os"
)

const (
	Prod = iota
	Devp
)

type LevelLogger interface {
	ChangeLevel(level int)
	Status(level int) abstract.StatusLogger
	Service(level int) abstract.ServiceLogger
	Database(level int) abstract.DatabaseLogger
}

type simpleLevelLogger struct {
	production bool
	status     abstract.StatusLogger
	service    abstract.ServiceLogger
	database   abstract.DatabaseLogger
}

func NewSimpleLevelLogger(level int) LevelLogger {
	logger := new(simpleLevelLogger)
	logger.ChangeLevel(level)
	logger.status = simple.NewSimpleStatusLogger(os.Stdout, log.Ldate|log.Ltime)
	logger.service = simple.NewSimpleServiceLogger(os.Stdout, log.Ldate|log.Ltime)
	logger.database = simple.NewSimpleDatabaseLogger(os.Stdout, log.Ldate|log.Ltime)
	return logger
}

func (mix *simpleLevelLogger) ChangeLevel(level int) {
	switch level {
	case Prod:
		mix.production = true
	default:
		mix.production = false
	}
}

func (mix *simpleLevelLogger) Status(level int) abstract.StatusLogger {
	if mix.production && level != Prod {
		return null.Status
	}
	return mix.status
}

func (mix *simpleLevelLogger) Service(level int) abstract.ServiceLogger {
	if mix.production && level != Prod {
		return null.Service
	}
	return mix.service
}

func (mix *simpleLevelLogger) Database(level int) abstract.DatabaseLogger {
	if mix.production && level != Prod {
		return null.Database
	}
	return mix.database
}
