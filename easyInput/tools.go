package easyInput

import "strings"

// CleanInput gets rid of the rubbish from input
func CleanInput(raw string) (output string) {
	output = strings.TrimSpace(raw)
	return output
}
