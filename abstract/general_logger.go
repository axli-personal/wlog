package abstract

// This a general logger only contains the most basic functionality.
type GeneralLogger interface {
	Print(args ...interface{})
}
