package logger

import "fmt"

// PrintInfo print message in screen with INFO level
func PrintInfo(messages ...interface{}) {
	fmt.Print("[   Info    ]")
	for _, message := range messages {
		fmt.Print(message)
	}
}

// PrintWarning print message in screen with WARNING level
func PrintWarning(messages ...interface{}) {
	fmt.Print("[   Warn    ]")
	for _, message := range messages {
		fmt.Print(message)
	}
}

// PrintError print message in screen with ERROR level
func PrintError(messages ...interface{}) {
	fmt.Print("[   Error   ]")
	for _, message := range messages {
		fmt.Print(message)
	}
}

// PrintSuccess print message in screen with SUCCESS level
func PrintSuccess(messages ...interface{}) {
	fmt.Print("[  Success  ]")
	for _, message := range messages {
		fmt.Print(message)
	}
}
