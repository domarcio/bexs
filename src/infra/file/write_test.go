package file

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestAppend(t *testing.T) {
	tmpFile, _ := ioutil.TempFile(os.TempDir(), "test-")
	open, err := os.OpenFile(tmpFile.Name(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		t.Fatal("error on create file")
	}

	defer func() {
		os.Remove(tmpFile.Name())
		tmpFile.Close()
		open.Close()
	}()

	t.Run("successful", func(t *testing.T) {
		wr := NewWriteFile(open)
		err := wr.Append("SINGLETEST1\n")
		if err != nil {
			t.Errorf("expected nil, got: %s", err.Error())
		}
		err = wr.Append("SINGLETEST2\n")
		if err != nil {
			t.Errorf("expected nil, got: %s", err.Error())
		}

		// Check content file
		newOpen, _ := os.Open(tmpFile.Name())
		strs, _ := ioutil.ReadAll(newOpen)

		if !strings.Contains(string(strs), "SINGLETEST") {
			t.Error("empty file")
		}
	})
	t.Run("error", func(t *testing.T) {
		wr := NewWriteFile(open)
		err := wr.Append("")
		if err == nil {
			t.Error("expected err, got nil")
		}
	})
}
