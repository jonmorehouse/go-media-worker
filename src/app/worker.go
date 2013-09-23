package app

// import proper projects needed for application
import (

	"bootstrap"
	"github.com/mikespook/gearman-go/worker"
)

// worker should take in a list of functions to "work" on 
// pass in a slice 
func Worker(functions[]*Task, killSignal chan int) {

	// ensure that we have set up the correct function here etc
	bootstrap.Bootstrap()

	// now lets create our correct worker etc
	gearman := bootstrap.CreateWorker()
	
	// loop through each of the functions and print the corresponding values
	for _, element := range functions {

		gearman.AddFunc(element.TaskName, element.Task, worker.Immediately)		
	}
	
	// now lets create a go routine to call this particular element
	go gearman.Work()

	// initialize a go routine for stopping the worker once started
	go func() {

		// now lets create a for loop to control the actual kill signal etc
		for {
		
			// capture the input variable that is pushed
			i := <- killSignal
			
			// check to see what type of signal we have sent
			if i == 0 {

				// close the connection
				gearman.Close()
			}
		}
	}()
}

