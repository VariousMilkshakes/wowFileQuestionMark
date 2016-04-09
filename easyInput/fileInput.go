package easyInput

import (
	"fmt"
	"io/ioutil"

	"github.com/variousmilkshakes/wowFileQuestionMark/goT"
)

// File contains methods and info about working file
type File struct {
	PlainText string
	Bytes     []byte
}

func openFile(filePath string) (f File, err error) {
	defer goT.Rp()

	contents, err := ioutil.ReadFile(filePath)
	goT.Cep(err)

	f = File{
		string(contents),
		contents,
	}

	fmt.Println(contents)

	return f, err
}
