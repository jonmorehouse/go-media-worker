package bootstrap_test

import (
	
	. "launchpad.net/gocheck"
	. "bootstrap"
)

// lets make sure that we load the client properly etc
func (s *TestSuite) TestGearmanWorker(c *C) {
	
	worker := CreateWorker()

	// make sure that the worker we created is valid
	c.Assert(worker, Not(Equals), nil)

	// now make sure that a function exists on this object
	c.Assert(worker.Work, Not(Equals), nil) 
}

// initialize and ensure that we are properly running this 
func (s *TestSuite) TestGearmanClient(c *C) {

	// initialize client
	client := CreateClient()

	// assert that we actually have created a non-nil pointer etc
	c.Assert(client, Not(Equals), nil)

	// lets assert that this is a valid client that has the elements that we need out of it etc	
	c.Assert(client, Not(Equals), nil) 

}


