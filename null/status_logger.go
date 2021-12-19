package null

type statusLogger struct {
}

func (status *statusLogger) Print(args ...interface{}) {
}

func (status *statusLogger) Printf(format string, args ...interface{}) {
}

func (status *statusLogger) Fatal(args ...interface{}) {
}

func (status *statusLogger) Fatalf(format string, args ...interface{}) {
}
