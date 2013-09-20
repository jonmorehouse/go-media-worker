package utilities_test

// import packages for application
import (
	
	. "launchpad.net/gocheck"
	. "utilities"
	"testing"
	"fmt"
)

// hook up gocheck to use the go test runner
func Test(t * testing.T) { TestingT(t) }

// test that keys are generated properly
func TestGenerateS3Keys(t *testing.T) {
	


}

// generate s3 key
func (s *S) TestGenerateS3KeysAreUnique(t *testing.T) {

	var key1, key2 string

	// generate 
	key1 = GenerateS3Keys()
	key2 = GenerateS3Keys()
	
	// now lets make sure the two unique keys were 
	t.Assert(key1, Not(Equals), key2)

		
}




