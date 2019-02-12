package logger

import (
	"flag"
	"log"
	"os"
)

// Level is the type for the various log levels
type Level int

const (
	// Debug Level for debugging purposes
	Debug Level = 10
	// Info Level for standard output
	Info Level = 20
	// Warn Level to print warnings
	Warn Level = 30
	// Error Level to print errors (will be print to ErrorWriter)
	Error Level = 40
)

var _globalLevel = flag.String("log-level", "info", "sets the desired log level (\"debug\", \"info\", \"warn\", \"error\")")

func getLevel() Level {
	switch *_globalLevel {
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "debug":
		return Debug
	}

	return Info
}

// Logger represents the logging object
type Logger struct {
	logger   *log.Logger
	errorLog *log.Logger
	level    Level
}

// New creates a new logging object
func New(prefix string) (logger *Logger) {
	logger = new(Logger)
	logger.logger = log.New(os.Stdout, prefix, 0)
	logger.errorLog = log.New(os.Stderr, prefix, 0)
	logger.level = getLevel()
	return
}

// Debugf outputs debug messages
func (logger *Logger) Debugf(format string, v ...interface{}) {
	if Debug < logger.level {
		return
	}

	logger.logger.Printf(format, v...)
}

// Infof outputs info messages
func (logger *Logger) Infof(format string, v ...interface{}) {
	if Info < logger.level {
		return
	}

	logger.logger.Printf(format, v...)
}

// Warnf outputs warning messages
func (logger *Logger) Warnf(format string, v ...interface{}) {
	if Warn < logger.level {
		return
	}

	logger.logger.Printf(format, v...)
}

// Errorf outputs error messages
func (logger *Logger) Errorf(format string, v ...interface{}) {
	if Error < logger.level {
		return
	}

	logger.errorLog.Printf(format, v...)
}
