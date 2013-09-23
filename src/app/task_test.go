package app_test

import (

	. "launchpad.net/gocheck"
	. "app"
)

func (s * TestSuite) TestNewTask(c *C) {
	
	// initialize task
	task := NewTask("test::sample", testTask)

	// ensure that we have stored the parameters properly
	c.Assert(task.Task, Not(Equals), nil)
	c.Assert(task.TaskName, Equals, "test::sample")
}
