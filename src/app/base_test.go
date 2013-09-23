package app_test

// import proper libraries as needed
import (

	. "launchpad.net/gocheck"
	"testing"
	. "app"
	"bootstrap"
)

// initialize test suite using gocheck
func Test(t *testing.T) { TestingT(t)}

// now initialize testSuite to use gocheck properly
type TestSuite struct {}
var _ = Suite(&TestSuite{})
var Tasks [1]Task

// initialize a testTask to be used through this package
func testTask () string {
	
	return "testTask"

}

// initialize a testSetupFunction here
func (s *TestSuite) SetUpSuite(c *C) {

	// initialize the bootstrap configuration 
	bootstrap.Bootstrap()
	
	Tasks[0] =  NewTask("test::sample", testTask)
}



