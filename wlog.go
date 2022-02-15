package wlog

type Options []struct {
	Key string
	Val interface{}
}

type Logger interface {
	// Log will output preprocessed options.
	//
	// Do not pass in any columns except you really know it.
	Log(options Options, columns ...string)

	// MakeHeaders will output headers.
	//
	// Pass in nil to generate headers.
	MakeHeaders(headers []string)
}
