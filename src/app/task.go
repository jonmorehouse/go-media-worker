package app

import (

	"github.com/mikespook/gearman-go/worker"	
)

// initialize a struct for containe
type Task struct {

	TaskName string	
	Task func(*worker.Job) ([]byte, error)

	// can put in other information about the run type / time of this application
}

// initialize an init function for creating new tasks
func NewTask(taskName string, taskFunction func(*worker.Job) ([]byte, error)) * Task {
	
	// initialize the actual struct
	task := new(Task)

	//set the taskName and the actual task
	task.TaskName = taskName
	task.Task = taskFunction	

	// return the newly created task
	return task
}
