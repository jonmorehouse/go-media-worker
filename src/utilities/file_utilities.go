package utilities

import "code.google.com/p/go-uuid/uuid"

func GenerateS3Key() string  {

	return uuid.New()

}
