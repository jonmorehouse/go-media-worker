package utilities

import (
	"code.google.com/p/go-uuid/uuid"
	"strings"
	"path"
)

// generate an s3Key
func GenerateS3Key(filePath string) string  {

	// slice the extension array and to grab just the last element, then grab the first index of that new piece
	extension := strings.Split(filePath, ".")[1:][0]
	
	// generate an array of the pieces we want to use in the fileKey
	pieces := []string{uuid.New(), extension}

	// now join those pieces with a . and return the string
	return strings.Join(pieces, ".")
}

// generate a unique output path to the directory passed in
func GenerateOutputPath(inputPath string, outputPath string) string {

	// grab the extension of the element	
	extension := strings.Split(inputPath, ".")[1:][0]

	// now lets initialize the various pieces of the application
	fileNamePieces := []string{uuid.New(), extension}

	return path.Join(outputPath, strings.Join(fileNamePieces, "."))

}
