package file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteToFileNonExisting(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "result.env")
	expectedVersion := "1.0.0"
	err := WriteToFile(expectedVersion, filename)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if string(data) != fmt.Sprintf("BUILD_VERSION=%s", expectedVersion) {
		t.Fail()
	}
}

func TestWriteToFileExisting(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "result.env")
	err := os.WriteFile(filename, []byte("test"), 0666)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	expectedVersion := "1.0.0"
	err = WriteToFile(expectedVersion, filename)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if string(data) != fmt.Sprintf("BUILD_VERSION=%s", expectedVersion) {
		t.Fail()
	}
}
