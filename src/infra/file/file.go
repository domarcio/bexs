package file

import (
	"errors"
	"os"
)

// NewCSVManager create a new manager to read and write files
func NewCSVManager(filename string) (Writer, Reader, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, nil, err
	}

	if info.IsDir() {
		return nil, nil, errors.New("is a directory")
	}

	read, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	write, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return nil, nil, err
	}

	return NewWriteFile(write), NewCSVfile(read, ","), nil
}
