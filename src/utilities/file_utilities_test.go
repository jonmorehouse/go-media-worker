package utilities_test

// import packages for application
import (
	
	. "launchpad.net/gocheck"
	. "utilities"
	"testing"
	"reflect"
	"strings"
)

// hook up gocheck to use the go test runner
func Test(t * testing.T) { TestingT(t) }

// now initialize a testSuite structure that will be passed into various method calls etc
type TestSuite struct {}
var _ = Suite(&TestSuite{})

// initialize some test strings that are used through the application
var testData = map[string]string {

	"sampleFilePath" : "test.MOV",
}


// set up test suite here 
func (s *TestSuite) SetUpSuite(c *C) {
	
	testData["sampleFilePathExtension"] = strings.Split(testData["sampleFilePath"], ".")[1:][0] 	

}

// test that keys are generated properly
// http://golang.org/doc/articles/laws_of_reflection.html
func (s *TestSuite) TestGenerateS3Key(c *C) {
	
	// generate a key to evaluate on
	key := GenerateS3Key(testData["sampleFilePath"])
	
	// grab the key reflection value through its interface
	keyType := reflect.TypeOf(key)

	// now make sure that the key type is a string
	c.Assert(keyType.Kind(), Equals, reflect.String)
}

// now make sure that the correct filenames are reflected within the generated keys
func (s * TestSuite) TestGenerateS3KeyPreservesExtension (c *C) {

	// initialize our fixure key element
	key := GenerateS3Key(testData["sampleFilePath"])

	// generate the keyExtension
	keyExtension := strings.Split(key, ".")[1:][0]

	// now lets make sure that the extensions are correct
	c.Assert(testData["sampleFilePathExtension"], Equals, keyExtension)
	
}

// make sure that generated keys are different
func (s *TestSuite) TestGenerateS3KeysAreUnique(c *C) {

	var key1, key2 string

	// generate 
	key1 = GenerateS3Key(testData["sampleFilePath"])
	key2 = GenerateS3Key(testData["sampleFilePath"])
	
	// now lets make sure the two unique keys were 
	c.Assert(key1, Not(Equals), key2)
}


