package database

const NULL = nullLogger(0)

type nullLogger int

func (database nullLogger) Print(operation string, table string, args ...interface{}) {
}

func (database nullLogger) Printf(operation string, table string, format string, args ...interface{}) {
}

func (database nullLogger) Fatal(operation string, table string, args ...interface{}) {
}

func (database nullLogger) Fatalf(operation string, table string, format string, args ...interface{}) {
}
