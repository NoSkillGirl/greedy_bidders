package log

import (
	"io"
	"log"
)

var (
	// Debug - only for development purposes
	Debug *log.Logger
	// Info - informational message
	Info *log.Logger
	// Warning - potential errors in coming versions
	Warning *log.Logger
	// Error - Immediate action to be taken
	Error *log.Logger
)

// SetupLogger - sets up different types of loggers
func SetupLogger(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Debug = log.New(traceHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
