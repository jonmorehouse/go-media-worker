package bootstrap_test

// import packages as needed for this application
import (

	. "launchpad.net/gocheck"
	"testing"
	"bootstrap"
	"reflect"
	//"github.com/bitly/go-simplejson"
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

// make sure that we initialize the config and it is valid with a key we can guarantee the existence of
func (s *TestSuite) TestBootstrapInitializesConfig(c *C) {

	// make sure that the Config pointer is empty here
	bootstrap.Bootstrap()
	
	// now lets assert that one of the various keys is not null
	// in the future, manually parse the config file and then run from there
	configPath, err := bootstrap.Config.Get("configPath").String()

	// make sure that no error gets returned
	c.Assert(err, Equals, nil)
	
	// now lets make sure that we get the proper name
	c.Assert(configPath, Not(Equals), "")
	c.Assert(reflect.TypeOf(configPath).Kind(), Equals, reflect.String)
}

// make sure that we don'r reload the configuration file multiple times if we happend to bootstrap() multiple timesx
func (s *TestSuite) TestBootstrapCalledOnlyOnce(c *C) {

	// make sure that the memory location of the pointer is not different for each call to bootstrap
	initialConfig := bootstrap.Config 

	// call the bootstrap function several times
	for i := 0; i < 10; i++ {
	
		bootstrap.Bootstrap()
		c.Assert(bootstrap.Config, Equals, initialConfig)
	}

	// assert that we still haven't called  the bootstrap function again
	c.Assert(bootstrap.Config, Equals, initialConfig)
}









