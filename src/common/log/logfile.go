package log

import (
	"fmt"
	"log"
	"os"
)

// Logfile struct
type Logfile struct{}

// NewLogfile create a new simple log
func NewLogfile(filename string, prefix string, flags int) *Logfile {
	if prefix != "" {
		log.SetPrefix(prefix)
	}

	if flags > 0 {
		log.SetFlags(flags)
	}

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("An error occurred on try open log file: %s", err.Error()))
	}

	log.SetOutput(logfile)

	return &Logfile{}
}

// Info log
func (l *Logfile) Info(message string, args ...interface{}) {
	message = fmt.Sprintf("[%s] %s", "INFO", message)
	log.Printf(message, args...)
}

// Warning log
func (l *Logfile) Warning(message string, args ...interface{}) {
	message = fmt.Sprintf("[%s] %s", "WARNING", message)
	log.Printf(message, args...)
}

// Error log
func (l *Logfile) Error(message string, args ...interface{}) {
	message = fmt.Sprintf("[%s] %s", "ERROR", message)
	log.Printf(message, args...)
}
