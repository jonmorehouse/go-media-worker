package app

import (

	"bootstrap"
)

// initialize a struct for containe
type Task struct {

	TaskName string	
	Task func()

	// can put in other information about the 
}

// initialize an init function for creating new tasks
func NewTask(taskName string, taskFunction func()) * Task {
	
	// initialize the actual struct
	task := new(Task)
	//set the taskName and the actual task
	task.TaskName = taskName
	task.Task = taskFunction	

	return task
}

// worker should take in a list of functions to "work" on 
func Worker(workerFunctions Task) {

	// ensure that we have set up the correct function here etc
	bootstrap.Bootstrap()

}


