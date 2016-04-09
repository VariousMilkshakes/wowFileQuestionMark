package exchange

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Settings holds keys and values
var Settings map[string]string

// ReadPhrase reads contact files
func ReadPhrase() map[string]string {
	input, err := ioutil.ReadFile("settings.ini")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	content := string(input)
	fmt.Println(content)

	Settings = make(map[string]string)

	lines := strings.Split(content, ",")
	for _, line := range lines {
		if checkForSetting(line) {
			setting := strings.Split(line, "=")
			Settings[setting[0]] = setting[1]
		}
	}

	return Settings
}

func checkForSetting(line string) bool { // Change to boolean result<<<<<<<<
	indx := strings.IndexRune(line, '#')
	fmt.Println(indx)
	if indx < 0 || indx > 5 {
		return true
	}

	return false
}
