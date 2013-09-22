package bootstrap

import (
	
	"log"

)

// do something clever in here in the future 
func ErrorMessage(err string) {

	log.Println(err)
}

// log a fatal system error
// calls os.Exit1 first
func FatalError(err error) {

	log.Fatalln(err)

}

// responsible for sending an error to the api we are working with
func SendError(err error) {
	
	log.Println(err)
	

}
