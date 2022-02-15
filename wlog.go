package wlog

type Pairs []struct {
	Key string
	Val interface{}
}

type Logger interface {
	// Log will output structured pairs.
	//
	// Don't pass in any columns unless you know the consequences.
	Log(pairs Pairs, columns ...string)

	// OutputHeaders will output headers.
	//
	// Don't pass in any headers unless you know the consequences.
	OutputHeaders(headers ...string)

	// Next will return next logger or nil.
	Next() Logger
}
