package app

import (
	
	"bootstrap"
	"github.com/mikespook/gearman-go/client"
	"sync"
)

// create a client that will send a channel back to recieve byte data
func PerformTask(taskName string, data[]byte) chan[]byte {
	
	// lets now create a client
	push := make(chan[]byte)

	// initialize a go function for running the client task etc
	go func() {

		// lets create a waitgroup etc
		var waitGroup sync.WaitGroup

		// now lets ensure that we add to this wait group -- ensure that execution occurs long enough
		waitGroup.Add(1)

		// initialize our client
		gearmanClient := bootstrap.CreateClient()

		// close the connection after this function ends
		defer gearmanClient.Close()

		// now initialize a jobHandler function for this element
		jobHandler := func(job * client.Job) {

			// push the job return to the calling function
			push <- job.Data	

			// we can now kill this go-routine
			waitGroup.Done()
		}

		// initialize the task and submit it to the client / server
		gearmanClient.Do(taskName, data, client.JOB_HIGH, jobHandler)

		// now lets wait for the application to be done 
		waitGroup.Wait()

	}()
	
	return push
}
