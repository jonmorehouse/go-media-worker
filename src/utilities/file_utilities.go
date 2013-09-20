package utilities

import (
	"code.google.com/p/go-uuid/uuid"
	"strings"
)

func GenerateS3Key(filePath string) string  {

	// slice the extension array and to grab just the last element, then grab the first index of that new piece
	extension := strings.Split(filePath, ".")[1:][0]
	
	// generate an array of the pieces we want to use in the fileKey
	pieces := []string{uuid.New(), extension}

	// now join those pieces with a . and return the string
	return strings.Join(pieces, ".")
}
