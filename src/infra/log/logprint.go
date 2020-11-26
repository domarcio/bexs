package log

import (
	"fmt"
)

// Logprint struct
type Logprint struct{}

// NewLogprint create a new simple log
func NewLogprint() *Logprint {
	return &Logprint{}
}

// Info log
func (*Logprint) Info(message string, args ...interface{}) {
	message = fmt.Sprintf("[%s] %s\n", "INFO", message)
	fmt.Printf(message, args...)
}

// Warning log
func (*Logprint) Warning(message string, args ...interface{}) {
	message = fmt.Sprintf("[%s] %s\n", "WARNING", message)
	fmt.Printf(message, args...)
}

// Error log
func (*Logprint) Error(message string, args ...interface{}) {
	message = fmt.Sprintf("[%s] %s\n", "ERROR", message)
	fmt.Printf(message, args...)
}
