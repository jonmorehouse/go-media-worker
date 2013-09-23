package encoder_test

// import proper libraries as needed
import (

	. "launchpad.net/gocheck"
	"testing"
	"bootstrap"
)

// initialize test suite using gocheck
func Test(t *testing.T) { TestingT(t)}

// now initialize testSuite to use gocheck properly
type TestSuite struct {}
var _ = Suite(&TestSuite{})

// initialize a testSetupFunction here
func (s *TestSuite) SetUpSuite(c *C) {

	// initialize the bootstrap configuration 
	bootstrap.Bootstrap()
}




