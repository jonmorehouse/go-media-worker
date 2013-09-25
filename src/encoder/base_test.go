package encoder_test

// import proper libraries as needed
import (

	. "launchpad.net/gocheck"
	"testing"
	"bootstrap"
	. "encoder"
	"path"
	//"fmt"
)

// initialize test suite using gocheck
func Test(t *testing.T) { TestingT(t)}

// now initialize testSuite to use gocheck properly
type TestSuite struct {}
var _ = Suite(&TestSuite{})

// lets create a sampleJob to be used for various tests
var FixturesDir string
var SampleClip string
var SampleJob *Job

// initialize a testSetupFunction here
func (s *TestSuite) SetUpSuite(c *C) {
		
	// initialize the bootstrap configuration 
	bootstrap.Bootstrap()

	// cache a pointer to the config
	config := bootstrap.Config

	// initialize and cache our FixturesDir
	FixturesDir = path.Join(config.Get("baseDir").MustString(), config.Get("fixturesDir").MustString())

	// now initialize the sampleClip 
	SampleClip = path.Join(FixturesDir, config.Get("fixtures").Get("clip").MustString())

	// initialize encodingType
	encodingType := config.Get("encodingTypes").MustArray()[0].(string)

	// now create a sampleJob
	SampleJob = NewJob(SampleClip, encodingType)
}


