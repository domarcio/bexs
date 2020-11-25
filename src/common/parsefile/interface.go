package parsefile

// Reader interface
type Reader interface {
	Current() []string
	Key() int
	Next()
	Rewind()
	Valid() bool
	Error() error
}
