package encoder

import (

	"os/exec"
	"bytes"
)

// return a channel which will return the status
func RunCommand(program string, arguments ...string) (chan bool, chan string) {

	// initialize a signal channel. This will be triggered if there is an error etc
	signal := make(chan bool, 1)

	// now lets make our string channel to send the output
	output := make(chan string)
	
	// this is the go function that is going to run the subcommand and will handle the error checking etc
	go func() {
	
		// pass the arguments that were passed to us to the function being called etc
		command := exec.Command(program, arguments...)

		// now lets go ahead and create an output buffer 
		var out bytes.Buffer
		command.Stdout = &out

		// now check for err
		err := command.Run()

		// if error is not valid then lets turn this off 
		if err != nil {
			
			// handle error here somehow
			signal <- false

		} else {
	
			// pipe this to the channel that is calling this element
			output <- out.String()
		}
	}() 

	// now return the various channels here etc
	return signal, output
}

