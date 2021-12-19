package code

// In most cases, user doesn't need to care about this.
// It will be useful when you get the exit code of the program.
const (
	StatusLoggerExit = iota + 2
	ServiceLoggerExit
	DatabaseLoggerExit
)
