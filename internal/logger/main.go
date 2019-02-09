package logger

import (
	"fmt"
	"io"
)

// VERBOSE describe VERBOSE log level
const VERBOSE uint = 2

// DEBUG describe DEBUG log level
const DEBUG uint = 3

// INFO describe INFO log level
const INFO uint = 4

// WARN describe WARN log level
const WARN uint = 5

// ERROR describe ERROR log level
const ERROR uint = 6

// Logger handle logging activity to the output target
type Logger struct {
	OutputTarget io.Writer
	LoggingLevel uint
}

// New creates instance of logger
func New(outputTarget io.Writer, loggingLevel uint) *Logger {
	return &Logger{
		OutputTarget: outputTarget,
		LoggingLevel: loggingLevel,
	}
}

func (logger *Logger) printfLevel(level uint, format string, messages ...interface{}) {
	if level >= logger.LoggingLevel {
		fmt.Fprintf(logger.OutputTarget, format, messages...)
	}
}

func (logger *Logger) printLevel(level uint, messages ...interface{}) {
	if level >= logger.LoggingLevel {
		fmt.Fprint(logger.OutputTarget, messages...)
	}
}

// PrintInfo print message in screen with INFO level
func (logger *Logger) PrintInfo(messages ...interface{}) {
	logger.printfLevel(INFO, "\033[0;34m[   Info    ] ")
	logger.printLevel(INFO, messages...)
	logger.printfLevel(INFO, "\033[0m\n")
}

// PrintWarning print message in screen with WARNING level
func (logger *Logger) PrintWarning(messages ...interface{}) {
	logger.printLevel(WARN, "\033[0;93m[   Warn    ] ")
	logger.printLevel(WARN, messages...)
	logger.printfLevel(WARN, "\033[0m\n")
}

// PrintError print message in screen with ERROR level
func (logger *Logger) PrintError(messages ...interface{}) {
	logger.printLevel(ERROR, "\033[0;31m[   Error   ] ")
	logger.printLevel(ERROR, messages...)
	logger.printfLevel(ERROR, "\033[0m\n")
}

// PrintSuccess print message in screen with SUCCESS level
func (logger *Logger) PrintSuccess(messages ...interface{}) {
	logger.printLevel(INFO, "\033[0;32m[  Success  ] ")
	logger.printLevel(INFO, messages...)
	logger.printfLevel(INFO, "\033[0m\n")
}

// Print print normal message in screen
func (logger *Logger) Print(level uint, messages ...interface{}) {
	logger.printLevel(level, messages...)
}

// Println print normal message in screen with endline
func (logger *Logger) Println(level uint, messages ...interface{}) {
	logger.printLevel(level, messages...)
	logger.printLevel(level, "\n")
}

// Printf print normal message in screen with format
func (logger *Logger) Printf(level uint, format string, messages ...interface{}) {
	logger.printfLevel(level, format, messages...)
}
