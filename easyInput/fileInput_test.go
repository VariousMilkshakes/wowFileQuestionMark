package easyInput

import (
	"fmt"
	"testing"
)

func TestFileOpen(t *testing.T) {
	file, err := openFile("flux-setup.exe")
	fmt.Println("E : ", err)
	teststring := CleanInput(file.PlainText)

	if teststring != "hello world" {
		t.Error("Expected 'hello world', got ", teststring)
	}

	t.Error("bytes : ", file.Bytes)
}
