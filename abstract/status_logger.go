package abstract

type StatusLogger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}