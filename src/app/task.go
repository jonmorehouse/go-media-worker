package app

// initialize a struct for containe
type Task struct {

	TaskName string	
	Task func()

	// can put in other information about the run type / time of this application
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
