package utilities_test

// import packages for application
import (
	
	. "launchpad.net/gocheck"
	. "utilities"
	"testing"
	"fmt"
	"reflect"
)

// hook up gocheck to use the go test runner
func Test(t * testing.T) { TestingT(t) }

// now initialize a testSuite structure that will be passed into various method calls etc
type TestSuite struct {}
var _ = Suite(&TestSuite{})

var testInt int = 22

// test that keys are generated properly
func (s *TestSuite) TestGenerateS3Keys(c *C) {
	
	// generate a key to evaluate on
	key := GenerateS3Key()
	
	// grab the key reflection value through its interface
	keyType := reflect.TypeOf(key)

	// now make sure that the key type is a string
	c.Assert(keyType.Kind(), Equals, reflect.String)
}

//// generate s3 key
//func (s *S) TestGenerateS3KeysAreUnique(t *testing.T) {

	//var key1, key2 string

	//// generate 
	//key1 = GenerateS3Keys()
	//key2 = GenerateS3Keys()
	
	//// now lets make sure the two unique keys were 
	//t.Assert(key1, Not(Equals), key2)

		
//}




