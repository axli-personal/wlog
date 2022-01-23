package status

const NULL = nullLogger(0)

type nullLogger int

func (status nullLogger) Print(args ...interface{}) {
}

func (status nullLogger) Printf(format string, args ...interface{}) {
}

func (status nullLogger) Fatal(args ...interface{}) {
}

func (status nullLogger) Fatalf(format string, args ...interface{}) {
}
