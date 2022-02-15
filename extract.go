package wlog

import (
	"fmt"
)

func WithExtract(logger Logger, keyword string) Logger {
	if logger == nil {
		panic("can't create logger from nil")
	}

	if len(keyword) == 0 {
		return logger
	}

	IncDepth(logger)

	return &columnExtractor{logger: logger, keyword: keyword}
}

type columnExtractor struct {
	logger  Logger
	keyword string
}

func (extractor *columnExtractor) Log(options Options, columns ...string) {
	column := "NULL"
	for i := 0; i < len(options); i++ {
		if options[i].Key == extractor.keyword {
			column = fmt.Sprint(options[i].Val)

			options[i].Key, options[i].Val = "", nil
			break
		}
	}

	columns = append(columns, column)
	extractor.logger.Log(options, columns...)
}

func (extractor *columnExtractor) MakeHeaders(headers []string) {
	headers = append(headers, extractor.keyword)
	extractor.logger.MakeHeaders(headers)
}
