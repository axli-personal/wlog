package mix

import (
	"github.com/axli-personal/wlog/database"
	"github.com/axli-personal/wlog/service"
	"github.com/axli-personal/wlog/status"
	"io"
)

const (
	Prod = iota
	Devp
)

type LevelLogger interface {
	ChangeLevel(level int)
	Status(level int) status.Logger
	Service(level int) service.Logger
	Database(level int) database.Logger
}

func NewLevelLogger(out io.Writer, flag int, level int) LevelLogger {
	logger := new(simpleLevelLogger)
	logger.ChangeLevel(level)
	logger.status = status.NewLogger(out, flag)
	logger.service = service.NewLogger(out, flag)
	logger.database = database.NewLogger(out, flag)
	return logger
}

type simpleLevelLogger struct {
	production bool
	status     status.Logger
	service    service.Logger
	database   database.Logger
}

func (mix *simpleLevelLogger) ChangeLevel(level int) {
	switch level {
	case Prod:
		mix.production = true
	default:
		mix.production = false
	}
}

func (mix *simpleLevelLogger) Status(level int) status.Logger {
	if mix.production && level != Prod {
		return status.NULL
	}
	return mix.status
}

func (mix *simpleLevelLogger) Service(level int) service.Logger {
	if mix.production && level != Prod {
		return service.NULL
	}
	return mix.service
}

func (mix *simpleLevelLogger) Database(level int) database.Logger {
	if mix.production && level != Prod {
		return database.NULL
	}
	return mix.database
}
