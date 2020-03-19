package log

import (
	"os"

	log "github.com/sirupsen/logrus"
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

var Log *log.Entry

// SetupLogger - sets up different types of loggers
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	Log = log.WithFields(log.Fields{
		// "common": "this is a common field",
		// "other":  "I also should be logged always",
	})
}
