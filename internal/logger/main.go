package logger

import "fmt"

// PrintInfo print message in screen with INFO level
func PrintInfo(messages ...interface{}) {
	fmt.Print("[   Info    ] ")
	fmt.Print(messages...)
	fmt.Println()
}

// PrintWarning print message in screen with WARNING level
func PrintWarning(messages ...interface{}) {
	fmt.Print("[   Warn    ] ")
	fmt.Print(messages...)
	fmt.Println()
}

// PrintError print message in screen with ERROR level
func PrintError(messages ...interface{}) {
	fmt.Print("[   Error   ] ")
	fmt.Print(messages...)
	fmt.Println()
}

// PrintSuccess print message in screen with SUCCESS level
func PrintSuccess(messages ...interface{}) {
	fmt.Print("[  Success  ] ")
	fmt.Print(messages...)
	fmt.Println()
}
