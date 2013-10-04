package encoder_test

import (

	. "launchpad.net/gocheck"
	. "encoder"
	"sync"
	"fmt"
)

// test out the run command element
func (s *TestSuite) TestExtractAudio(c *C) {

	//RunCommand([]string{"ls", "~/Documents"})
	ExtractAudio(SampleJob)

}

func (s *TestSuite) TestRunCommand(c *C) {
	
	// first set up a wait group for running etc
	var waitGroup sync.WaitGroup

	// add an element to the wait group
	waitGroup.Add(1)

	// now lets get the result channel etc
	signal, output := RunCommand("ls", []string{"/Users/MorehouseJ09/Documents"}...)

	// wait for the timeout
	select {

		// now listen for the signal to come in
		case ok, _ := <- signal:

			if ok {

				fmt.Println("Good")

			} else {

				fmt.Println("BAD")
			}

			waitGroup.Done()
			break

		case value, ok := <- output:

			if !ok {
				
				// handle error here 
			}

			fmt.Println(value)

			waitGroup.Done()
			break

	}

	// initialize element
	waitGroup.Wait()

}

func (s *TestSuite) TestAudioOnly(c *C) {

	//ExtractAudio(SampleJob)	

}
