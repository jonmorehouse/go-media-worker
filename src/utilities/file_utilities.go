package utilities

import "code.google.com/p/go-uuid/uuid"

func GenerateS3Keys() string  {

	return uuid.New()

}
