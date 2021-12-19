// This package implement three types of logger but they do nothing.
package null

var (
	Status   = new(statusLogger)
	Service  = new(serviceLogger)
	Database = new(databaseLogger)
)
