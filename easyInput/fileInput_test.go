package easyInput

import (
	"fmt"
	"testing"
)

func TestFileOpen(t *testing.T) {
	file, err := openFile("testFile.txt")
	fmt.Println("E : ", err)
	teststring := CleanInput(file.PlainText)

	if teststring != "hello world" {
		t.Error("Expected 'hello world', got ", teststring)
	}
}
