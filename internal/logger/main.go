package logger

import (
	"fmt"
	"os"
)

// PrintInfo print message in screen with INFO level
func PrintInfo(messages ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[0;34m[   Info    ] ")
	fmt.Fprint(os.Stderr, messages...)
	fmt.Fprintf(os.Stderr, "\033[0m\n")
}

// PrintWarning print message in screen with WARNING level
func PrintWarning(messages ...interface{}) {
	fmt.Fprint(os.Stderr, "\033[0;93m[   Warn    ] ")
	fmt.Fprint(os.Stderr, messages...)
	fmt.Fprintf(os.Stderr, "\033[0m\n")
}

// PrintError print message in screen with ERROR level
func PrintError(messages ...interface{}) {
	fmt.Fprint(os.Stderr, "\033[0;31m[   Error   ] ")
	fmt.Fprint(os.Stderr, messages...)
	fmt.Fprintf(os.Stderr, "\033[0m\n")
}

// PrintSuccess print message in screen with SUCCESS level
func PrintSuccess(messages ...interface{}) {
	fmt.Fprint(os.Stderr, "\033[0;32m[  Success  ] ")
	fmt.Fprint(os.Stderr, messages...)
	fmt.Fprintf(os.Stderr, "\033[0m\n")
}
