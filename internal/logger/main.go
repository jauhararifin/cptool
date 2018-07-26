package logger

import (
	"fmt"
	"io"
	"os"
)

// OutputTarget specify where log should be written
var OutputTarget io.Writer = os.Stderr

// PrintInfo print message in screen with INFO level
func PrintInfo(messages ...interface{}) {
	fmt.Fprintf(OutputTarget, "\033[0;34m[   Info    ] ")
	fmt.Fprint(OutputTarget, messages...)
	fmt.Fprintf(OutputTarget, "\033[0m\n")
}

// PrintWarning print message in screen with WARNING level
func PrintWarning(messages ...interface{}) {
	fmt.Fprint(OutputTarget, "\033[0;93m[   Warn    ] ")
	fmt.Fprint(OutputTarget, messages...)
	fmt.Fprintf(OutputTarget, "\033[0m\n")
}

// PrintError print message in screen with ERROR level
func PrintError(messages ...interface{}) {
	fmt.Fprint(OutputTarget, "\033[0;31m[   Error   ] ")
	fmt.Fprint(OutputTarget, messages...)
	fmt.Fprintf(OutputTarget, "\033[0m\n")
}

// PrintSuccess print message in screen with SUCCESS level
func PrintSuccess(messages ...interface{}) {
	fmt.Fprint(OutputTarget, "\033[0;32m[  Success  ] ")
	fmt.Fprint(OutputTarget, messages...)
	fmt.Fprintf(OutputTarget, "\033[0m\n")
}
