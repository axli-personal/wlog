package abstract

// Logging when perform database operation.
type DatabaseLogger interface {
	Print(operation string, table string, args ...interface{})
	Printf(operation string, table string, format string, args ...interface{})

	Fatal(operation string, table string, args ...interface{})
	Fatalf(operation string, table string, format string, args ...interface{})
}