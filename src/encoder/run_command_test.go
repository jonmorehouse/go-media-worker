package encoder_test

import (

	. "launchpad.net/gocheck"
	. "encoder"
	"sync"
)

func (s *TestSuite) TestRunCommand(c *C) {
	
	// first set up a wait group for running etc
	var waitGroup sync.WaitGroup

	// add an element to the wait group
	waitGroup.Add(1)

	// now lets get the result channel etc
	signal, output := RunCommand("ls", []string{"/Users/MorehouseJ09/Documents"}...)

	// wait for the timeout
	select {

		// now listen for the signal to come in
		case ok, _ := <- signal:

			// can finish up process
			c.Assert(ok, Equals, true)

			waitGroup.Done()
			break

		case value, ok := <- output:

			// now ensure that our status is okay
			c.Assert(ok, Equals, true)

			// ensure we get a valid string passed back to us
			c.Assert(value, Not(Equals), nil)

			// we can now finish up the waiting etc
			waitGroup.Done()
			break

	}

	// initialize element
	waitGroup.Wait()
}

