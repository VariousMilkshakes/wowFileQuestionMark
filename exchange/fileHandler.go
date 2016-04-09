package exchange

import (
	"regexp"
)

func fileNameFromPath(filePath string) (fileName string) {
	parts := regexp.MustCompile("[/\\]").Split(filePath, -1)
	fileName = parts[len(parts)-1]

	return fileName
}
