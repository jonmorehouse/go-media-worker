package bootstrap_test

// import packages as needed for this application
import (

	. "launchpad.net/gocheck"
	. "bootstrap"
	"testing"
	//"reflect"
)

// initialize test suite using gocheck
func Test(t *testing.T) { TestingT(t)}

// now initialize testSuite to use gocheck properly
type TestSuite struct {}
var _ = Suite(&TestSuite{})

func (s *TestSuite) TestBootstrap(c *C) {

	Bootstrap()
	Bootstrap()
	Bootstrap()

}






