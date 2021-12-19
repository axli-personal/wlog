package null

type serviceLogger struct {
}

func (service *serviceLogger) Print(url string, args ...interface{}) {
}

func (service *serviceLogger) Printf(url string, format string, args ...interface{}) {
}

func (service *serviceLogger) Fatal(url string, args ...interface{}) {
}

func (service *serviceLogger) Fatalf(url string, format string, args ...interface{}) {
}
