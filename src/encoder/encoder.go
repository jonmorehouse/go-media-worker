package encoder 

import (

	"os"
)

// initialize the encode job element
// this is a wrapper
func Encode(job * Job) {

	// ensure that the file exists
	if _, err := os.Stat(job.InputPath); err != nil {
		
		// initialize the various elements
		job.Status = "failed"
		job.Finished = false

		// emit some channel elements here for other go routines to be warned etc?
		return
	}


}


