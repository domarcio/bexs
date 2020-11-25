package file

// Reader interface
type Reader interface {
	Current() []string
	Key() int
	Next()
	Rewind()
	Valid() bool
	Error() error
	CloseFile() error
}

// Writer interface
type Writer interface {
	Append(txt string) error
	CloseFile() error
}
