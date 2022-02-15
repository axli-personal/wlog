package wlog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func NewLogger(writer io.Writer, split string) Logger {
	if writer == nil {
		writer = os.Stdout
	}
	if len(split) == 0 {
		split = "||"
	}
	return &tableLogger{out: writer, split: split}
}

type tableLogger struct {
	split string
	lock  sync.Mutex
	out   io.Writer
}

func (l *tableLogger) Log(options Options, columns ...string) {
	l.lock.Lock()

	for i := 0; i < len(columns); i++ {
		fmt.Fprint(l.out, columns[i], l.split)
	}

	last := len(options) - 1

	for i := 0; i < last; i++ {
		if len(options[i].Key) != 0 || options[i].Val != nil {
			fmt.Fprint(l.out, options[i].Key, "=", options[i].Val, ";")
		}
	}

	// Print the last pair.
	if last > -1 {
		if len(options[last].Key) != 0 || options[last].Val != nil {
			fmt.Fprint(l.out, options[last].Key, "=", options[last].Val)
		}
	}

	fmt.Fprint(l.out, "\n")

	l.lock.Unlock()
}

func (l *tableLogger) MakeHeaders(headers []string) {
	l.lock.Lock()

	for i := 0; i < len(headers); i++ {
		fmt.Fprint(l.out, headers[i], l.split)
	}

	fmt.Fprint(l.out, "Details\n")

	l.lock.Unlock()
}
