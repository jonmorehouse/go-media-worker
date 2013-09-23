package app_test

// initialize tests to test and verify the worker etc
import (
	. "launchpad.net/gocheck"
	. "app"
	"github.com/mikespook/gearman-go/client"
	"sync"
)


// test the goRoutine runner. This should be automatically started when we call the function etc
func (s *TestSuite) TestWorkerGoRoutine(c *C) {
	
	// create a channel for triggering our workers on / off
	push := make(chan int)

	// create a wait group to manage go routines running else where
	var waitGroup sync.WaitGroup

	// add to our wait group because we want to wait for the various tasks to complete etc 
	waitGroup.Add(1)
	
	// initialize our worker and create a worker goroutine for handling tasks
	Worker(Tasks[0:1], push)	
	
	// now we actually want to handle clients etc and ensure that we are getting a basic response back ...
	gearmanClient, _ := client.New("127.0.0.1:4730")

	// close the connection once we are finished
	defer gearmanClient.Close()

	// now create sample data to send to the element
	data := []byte("sampledata")

	// call back function for job
	jobHandler := func(job *client.Job) {

		c.Assert(string(job.Data), Equals, "test::sample::return") 

		// we can kill our server now that we know its working
		push <- 0

		// decrement the wait group now 
		waitGroup.Done()
	}

	// Initialize task etc 
	gearmanClient.Do(Tasks[0].TaskName, data, client.JOB_HIGH, jobHandler)

	// now wait for until we have a response from the server before doing anything else
	waitGroup.Wait()
}

