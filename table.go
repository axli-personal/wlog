package wlog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// NewLogger will create a basic logger.
//
// Default: (out == nil) -> os.Stdout.
//
// Default: (split == "") -> "||".
func NewLogger(out io.Writer, split string) Logger {
	if out == nil {
		out = os.Stdout
	}
	if len(split) == 0 {
		split = "||"
	}
	return &tableLogger{out: out, split: split}
}

type tableLogger struct {
	split string
	lock  sync.Mutex
	out   io.Writer
}

func (l *tableLogger) Log(pairs Pairs, columns ...string) {
	l.lock.Lock()

	for i := 0; i < len(columns); i++ {
		fmt.Fprint(l.out, columns[i], l.split)
	}

	last := len(pairs) - 1

	for last > -1 {
		if len(pairs[last].Key) != 0 || pairs[last].Val != nil {
			break
		}
		last--
	}

	for i := 0; i < last; i++ {
		if len(pairs[i].Key) != 0 || pairs[i].Val != nil {
			fmt.Fprint(l.out, pairs[i].Key, "=", pairs[i].Val, ";")
		}
	}

	// Print the last pair.
	if last > -1 {
		fmt.Fprint(l.out, pairs[last].Key, "=", pairs[last].Val)
	}

	fmt.Fprint(l.out, "\n")

	l.lock.Unlock()
}

func (l *tableLogger) OutputHeaders(headers ...string) {
	l.lock.Lock()

	for i := 0; i < len(headers); i++ {
		fmt.Fprint(l.out, headers[i], l.split)
	}

	fmt.Fprint(l.out, "Details\n")

	l.lock.Unlock()
}

func (l *tableLogger) Next() Logger {
	return nil
}
