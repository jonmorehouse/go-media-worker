package bootstrap_test

import (
	
	. "launchpad.net/gocheck"
	. "bootstrap"
	"reflect"
)

func (s *TestSuite) TestErrorMessage(c *C) {
	
	// make sure that we have introduced the correct functions into the current namespace
	c.Assert(reflect.TypeOf(ErrorMessage).Kind(), Equals, reflect.Func)
}

// test fatal errors and ensure that we are introducing the correct functions into the namespace
func (s *TestSuite) TestFatalError(c *C) {

	c.Assert(reflect.TypeOf(FatalError).Kind(), Equals, reflect.Func)

}

// this will be a complicated, rpc-wrapping error in the future etc
func (s *TestSuite) TestSendError(c *C) {

	c.Assert(reflect.TypeOf(SendError).Kind(), Equals, reflect.Func)

}
