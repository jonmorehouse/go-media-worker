package encoder

import (
	
	"utilities"
	"bootstrap"
	//"fmt"
)

// initailize a job struct for managing the paths and various pieces of the encoding task etc 
type Job struct {

	// initalize 
	InputPath string 
	// cache the outputDir in case we need to reset / restart the directory
	OutputDir string
	// initialize our full path to the output path
	OutputPath string 
	// initialize the encodingType
	EncodingType string

	//
	Status string

	Finished bool

	// initialize a channel for returning elements / statuses
	Channel chan[]byte
}

// create a new job with the inputPath 
func NewJob(inputPath string, encodingType string)  * Job {
		
	// initialize job and store a pointer to it
	job := new(Job)

	// initialize the various elements needed
	job.InputPath = inputPath
	job.OutputDir = bootstrap.Config.Get("outputDir").MustString()
	job.OutputPath = utilities.GenerateOutputPath(inputPath, job.OutputDir)
	job.EncodingType = encodingType
	
	// job.status = ["initialized", "processing", "error", "success"]
	job.Status = "initialized"

	job.Finished = false

	// now initialize the channel element
	job.Channel = make(chan[]byte)

	// 
	return job
}

// reset a job -- in case a job fails the first time around
func (self * Job) Reset() {
	
	// reset the various elements that need to be reset 
	self.Finished = false	
	self.Status = "initialized"
	self.OutputPath = utilities.GenerateOutputPath(self.OutputPath, self.OutputDir)

}

