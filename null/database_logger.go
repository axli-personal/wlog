package null

type databaseLogger struct {
}

func (database *databaseLogger) Print(operation string, table string, args ...interface{}) {
}

func (database *databaseLogger) Printf(operation string, table string, format string, args ...interface{}) {
}

func (database *databaseLogger) Fatal(operation string, table string, args ...interface{}) {
}

func (database *databaseLogger) Fatalf(operation string, table string, format string, args ...interface{}) {
}
