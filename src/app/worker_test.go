package app_test

// initialize tests to test and verify the worker etc
import (

	. "launchpad.net/gocheck"
	. "app"
)

// test the goRoutine runner. This should be automatically started when we call the function etc
func (s *TestSuite) TestWorkerGoRoutine(c *C) {
	
	Worker()	

}
