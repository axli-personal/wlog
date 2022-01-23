package service

const NULL = nullLogger(0)

type nullLogger int

func (service nullLogger) Print(url string, args ...interface{}) {
}

func (service nullLogger) Printf(url string, format string, args ...interface{}) {
}

func (service nullLogger) Fatal(url string, args ...interface{}) {
}

func (service nullLogger) Fatalf(url string, format string, args ...interface{}) {
}
