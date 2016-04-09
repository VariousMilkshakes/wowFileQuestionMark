package easyInput

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/variousmilkshakes/wowFileQuestionMark/goT"
)

// File contains methods and info about working file
type File struct {
	Bytes     []byte
	FileName  string
	Path      string
	PlainText string
}

// WriteFile writes file
func (file *File) WriteFile() {
	newFile, err := os.Create(file.FileName)
	goT.Cep(err)

	rep, err := newFile.Write(file.Bytes)
	goT.Cep(err)

	fmt.Println("Saved file ", file.FileName)
	fmt.Println(rep)

	defer newFile.Close()
}

// OpenFile returns file struct
func OpenFile(filePath string) (f File, err error) {
	defer goT.Rp()

	contents, err := ioutil.ReadFile(filePath)
	goT.Cep(err)

	f = File{
		contents,
		fileNameFromPath(filePath),
		filePath,
		string(contents),
	}

	fmt.Println(contents)

	return f, err
}

func fileNameFromPath(filePath string) (fileName string) {
	parts := regexp.MustCompile("[\\/]").Split(filePath, -1)
	fileName = parts[len(parts)-1]

	return fileName
}
