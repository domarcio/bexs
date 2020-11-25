package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

var dummyCSVContent = []string{
	"FOO,BAR",
	"BAR,FOO",
	"OOF,RAB",
}

func createTempFile() *os.File {
	tmpFile, _ := ioutil.TempFile(os.TempDir(), "test-")

	for _, l := range dummyCSVContent {
		l += "\n"
		text := []byte(l)
		tmpFile.Write(text)
	}

	open, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	return open
}

func TestIterator(t *testing.T) {
	tmpFile := createTempFile()
	defer func() {
		os.Remove(tmpFile.Name())
		tmpFile.Close()
	}()

	csvreader := NewCSVfile(tmpFile, ",")

	assertTests := []struct {
		arg1    string
		arg2    string
		key     int
		isValid bool
	}{
		{
			arg1:    "FOO",
			arg2:    "BAR",
			key:     1,
			isValid: true,
		},
		{
			arg1:    "BAR",
			arg2:    "FOO",
			key:     2,
			isValid: true,
		},
		{
			arg1:    "OOF",
			arg2:    "RAB",
			key:     3,
			isValid: true,
		},
	}

	for _, assert := range assertTests {
		tk := fmt.Sprintf("%s_%s", assert.arg1, assert.arg2)
		t.Run(tk, func(t *testing.T) {
			csvreader.Next()

			line := csvreader.Current()
			if len(line) <= 0 {
				t.Error("expected a line")
			}

			first := line[0]
			second := line[1]
			key := csvreader.Key()
			isValid := csvreader.Valid()
			firstExp := assert.arg1
			secondExp := assert.arg2
			keyExp := assert.key
			isValidExp := assert.isValid

			if first != firstExp {
				t.Errorf("expected found %s, got %s", firstExp, first)
			}
			if second != secondExp {
				t.Errorf("expected found %s, got %s", secondExp, second)
			}
			if isValid != isValidExp {
				t.Errorf("expected found %+v, got %+v", isValidExp, isValid)
			}
			if key != keyExp {
				t.Errorf("expected found %d, got %d", keyExp, key)
			}
		})
	}

	t.Run("eof", func(t *testing.T) {
		csvreader.Next()
		if len(csvreader.Current()) > 0 {
			t.Error("expected an empry line")
		}
		if csvreader.Valid() {
			t.Error("expected a invalid iterator")
		}
		err := csvreader.Error()
		if err != io.EOF {
			if err != nil {
				t.Errorf("expected an io.EOF error, got %s", err.Error())
			} else {
				t.Errorf("expected an io.EOF error")
			}
		}
	})

	csvreader.Rewind()
	for _, assert := range assertTests {
		tk := fmt.Sprintf("rewind_%s_%s", assert.arg1, assert.arg2)
		t.Run(tk, func(t *testing.T) {
			csvreader.Next()

			line := csvreader.Current()
			if len(line) <= 0 {
				t.Error("expected a line")
			}

			first := line[0]
			second := line[1]
			key := csvreader.Key()
			isValid := csvreader.Valid()
			firstExp := assert.arg1
			secondExp := assert.arg2
			keyExp := assert.key
			isValidExp := assert.isValid

			if first != firstExp {
				t.Errorf("expected found %s, got %s", firstExp, first)
			}
			if second != secondExp {
				t.Errorf("expected found %s, got %s", secondExp, second)
			}
			if isValid != isValidExp {
				t.Errorf("expected found %+v, got %+v", isValidExp, isValid)
			}
			if key != keyExp {
				t.Errorf("expected found %d, got %d", keyExp, key)
			}
		})
	}
}
