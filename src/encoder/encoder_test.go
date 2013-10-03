package encoder_test

import (

	. "launchpad.net/gocheck"
	. "encoder"
)

func (s *TestSuite) TestClient(c *C) {

	// run the encoding job 
	Encode(SampleJob)
	
	// now run some tests to ensure that things are working properly etc

	
}



