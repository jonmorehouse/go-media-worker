package bootstrap_test

import (
	
	. "launchpad.net/gocheck"
	. "bootstrap"
	"fmt"
)

func (s *TestSuite) TestFatalErrorMessage(c *C) {

	fmt.Println("TESTING 123")	
	ErrorMessage("TEST")
}
