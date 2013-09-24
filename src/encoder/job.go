package encoder

import (
	
	//"utilities"
	"fmt"
)

type Job struct {

	// initalize 
	InputPath string 
	OutputPath string 
	EncodingType string
	Status string

	// initialize a channel for returning elements / statuses
	Channel chan[]byte
}

// create a new job with the inputPath 
func NewJob(inputPath string, encodingType string)  * Job {
		
	// initialize job and store a pointer to it
	job := new(Job)

	// initialize the various elements needed
	job.InputPath = inputPath
	job.EncodingType = encodingType
	//job.OutputPath =

	return job
}

func (self * Job) Reset() {
	
	fmt.Println("TEST RESET")


}
