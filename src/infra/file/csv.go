package file

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// CSVFile struct
type CSVFile struct {
	reader      *bufio.Scanner
	osfile      *os.File
	split       string
	currentLine []string
	currentErr  error
	currentKey  int
}

// NewCSVfile create a new csv file iterator
func NewCSVfile(osfile *os.File, split string) *CSVFile {
	scan := bufio.NewScanner(osfile)
	scan.Split(bufio.ScanLines)

	if split == "" {
		split = ","
	}

	return &CSVFile{reader: scan, split: split, osfile: osfile}
}

// Current line
func (c *CSVFile) Current() []string {
	return c.currentLine
}

// Key current line
func (c *CSVFile) Key() int {
	return c.currentKey
}

// Next move to the next line
func (c *CSVFile) Next() {
	c.currentLine = []string{}

	result := c.reader.Scan()
	line := c.reader.Text()
	err := c.reader.Err()

	if err != nil {
		c.currentErr = err
		return
	}

	if !result {
		c.currentErr = io.EOF
		return
	}

	c.currentLine = strings.Split(line, c.split)
	c.currentKey++
}

// Rewind go back to the first line
func (c *CSVFile) Rewind() {
	_, err := c.osfile.Seek(io.SeekStart, io.SeekStart)
	if err != nil {
		c.currentErr = err
		return
	}

	scan := bufio.NewScanner(c.osfile)
	scan.Split(bufio.ScanLines)

	c.currentKey = 0
	c.reader = scan
	c.currentErr = nil
}

// Valid check if is a valid line
func (c *CSVFile) Valid() bool {
	return c.currentErr == nil
}

// Error returns a current error line
func (c *CSVFile) Error() error {
	return c.currentErr
}

// CloseFile osfile
func (c *CSVFile) CloseFile() error {
	return c.osfile.Close()
}
