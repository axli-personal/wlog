// Logging system designed for backend web application.
package wlog

type Options []struct {
	Key string
	Val interface{}
}

type Logger interface {
	Log(options Options, columns ...string)
	MakeHeaders(headers []string)
}
