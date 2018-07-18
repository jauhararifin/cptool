package logger

import "fmt"

// PrintInfo print message in screen with INFO level
func PrintInfo(messages ...interface{}) {
	fmt.Printf("\033[0;34m[   Info    ] ")
	fmt.Print(messages...)
	fmt.Printf("\033[0m\n")
}

// PrintWarning print message in screen with WARNING level
func PrintWarning(messages ...interface{}) {
	fmt.Print("\033[0;93m[   Warn    ] ")
	fmt.Print(messages...)
	fmt.Printf("\033[0m\n")
}

// PrintError print message in screen with ERROR level
func PrintError(messages ...interface{}) {
	fmt.Print("\033[0;31m[   Error   ] ")
	fmt.Print(messages...)
	fmt.Printf("\033[0m\n")
}

// PrintSuccess print message in screen with SUCCESS level
func PrintSuccess(messages ...interface{}) {
	fmt.Print("\033[0;32m[  Success  ] ")
	fmt.Print(messages...)
	fmt.Printf("\033[0m\n")
}
