package encoder_test

import (

	. "launchpad.net/gocheck"
	. "encoder"
)

// test out the run command element
func (s *TestSuite) TestExtractAudio(c *C) {

	//RunCommand([]string{"ls", "~/Documents"})
	ExtractAudio(SampleJob)

}

func (s *TestSuite) TestAudioOnly(c *C) {

	//ExtractAudio(SampleJob)	

}
