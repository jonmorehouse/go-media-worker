package encoder_test

// import proper libraries as needed
import (

	. "launchpad.net/gocheck"
	"testing"
	"bootstrap"
	. "encoder"
	"fmt"
	"path"

)

// initialize test suite using gocheck
func Test(t *testing.T) { TestingT(t)}

// now initialize testSuite to use gocheck properly
type TestSuite struct {}
var _ = Suite(&TestSuite{})

// lets create a sampleJob to be used for various tests
var FixturesDir string
var SampleJob Job

// initialize a testSetupFunction here
func (s *TestSuite) SetUpSuite(c *C) {

	// initialize the bootstrap configuration 
	bootstrap.Bootstrap()
	config := bootstrap.Config

	FixturesDir = path.Join(config.Get("baseDir").MustString(), config.Get("fixturesDir").MustString())

	fmt.Println(FixturesDir)
}




