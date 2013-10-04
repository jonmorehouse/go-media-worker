package encoder

import (

	//"fmt"
	//"os/exec"
	//"bytes"
)

// return a channel which will return the status
func RunCommand(arguments []string) (chan bool, chan string) {

	// initialize a signal channel. This will be triggered if there is an error etc
	signal := make(chan bool, 1)

	// now lets make our string channel to send the output
	output := make(chan string)

	// this is the go function that is going to run the subcommand and will handle the error checking etc
	go func() {
	
		// run the command and parse results. We need to notify the element of error etc		
		// signalling false is not good!
		signal <- false
		
	

	}() 

	// now return the various channels here etc
	return signal, output
}

func ExtractAudio(job *Job) {

/*
	// generate the command and then call it on the command line
	app := "ffmpeg"
	outputPath := ""

	// now lets actually run this element etc	
	cmd := exec.Command(app, outputPath) 

	// initialize output elements etc
	output, err := cmd.Output()

	// initialize error etc
	if err != nil {
		
		//fmt.println(err.Error())
		//fmt.Println(err.Error())
		return
	}

	//fmt.Println(string(output))
*/

}

func encodeVideo(job *Job) {


}
