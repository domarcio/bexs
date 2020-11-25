package file

import (
	"errors"
	"os"
)

// WriteFile struct
type WriteFile struct {
	osfile *os.File
}

// NewWriteFile create a new writer file
func NewWriteFile(osfile *os.File) *WriteFile {
	return &WriteFile{osfile: osfile}
}

// Append append to eof
func (w *WriteFile) Append(txt string) error {
	n, err := w.osfile.WriteString(txt)
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("no text written in the file")
	}
	return nil
}

// CloseFile osfile
func (w *WriteFile) CloseFile() error {
	return w.osfile.Close()
}
