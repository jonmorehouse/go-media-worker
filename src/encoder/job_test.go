package encoder_test

// import correct libraries for testing application
import (

	. "launchpad.net/gocheck"
	"bootstrap"
	//"testing"
	"fmt"
	//. "encoder"
)

// now create a basic element etc
func (s *TestSuite) TestEncodingJob(c *C) {
	
	fmt.Println(bootstrap.Config.Get("gearmanHosts"))
	//fmt.Println(bootstrap.Config.Get("fixtures").Get("inputPath"))

}






