package log

// Logger log interface
type Logger interface {
	Info(message string, args ...interface{})
	Warning(message string, args ...interface{})
	Error(message string, args ...interface{})
}
