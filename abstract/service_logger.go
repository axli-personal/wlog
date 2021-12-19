package abstract

// Logging in specific service.
type ServiceLogger interface {
	Print(url string, args ...interface{})
	Printf(url string, format string, args ...interface{})

	Fatal(url string, args ...interface{})
	Fatalf(url string, format string, args ...interface{})
}