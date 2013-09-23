package app_test


// initialize tests to test and verify the worker etc
import (
	. "launchpad.net/gocheck"
	. "app"
	"fmt"
)

func (s *TestSuite) TestClient(c *C) {

	// create a channel so that we can stop the worker later on
	workerPush := make(chan int)
	
	// create our worker function / go-routine
	Worker(Tasks[0:1], workerPush)

	// initialize a channel that will listen for the client's response
	jobPull := PerformTask(Tasks[0].TaskName, []byte(""))
	
	// now lets just patiently wait for the jobResults that are being run through the gearmanServers
	for {
		// grab the jobResults
		jobResults := <- jobPull

		// now if we actually have jobResults then go ahead and work some magic!
		if jobResults != nil {

			// kill the server
			workerPush <- 0	

			// ensure that we get the correct response from the server
			c.Assert(string(jobResults), Equals, "test::sample::return")

			// now we can break from this task
			break
		}
	}
}



