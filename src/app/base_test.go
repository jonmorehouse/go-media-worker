package app_test

// import proper libraries as needed
import (

	. "launchpad.net/gocheck"
	"testing"
	. "app"
	"bootstrap"
	"github.com/mikespook/gearman-go/worker"
)

// initialize test suite using gocheck
func Test(t *testing.T) { TestingT(t)}

// now initialize testSuite to use gocheck properly
type TestSuite struct {}
var _ = Suite(&TestSuite{})

// store a list of pointers to the various tasks that we want to set up in this application
var Tasks [1]*Task

// initialize a testTask to be used through this package
func testTask (job *worker.Job) ([]byte, error) {

	// return a sample array for now
	returnArray := []byte("test::sample::return")

	return returnArray, nil
}

// initialize a testSetupFunction here
func (s *TestSuite) SetUpSuite(c *C) {

	// initialize the bootstrap configuration 
	bootstrap.Bootstrap()
	
	// initialize Task for this application
	Tasks[0] = NewTask("test::sample", testTask)
}



