package commands

import (
	"fmt"

	"github.com/gaurav4288/4_task_tracker/internal/task"
)

func ListTasks(args ...string) error {
	store, err := task.Load()
	if err != nil {
		return err
	}

	var tasks []task.Task
	if len(args) == 0 {
		tasks = store.Tasks
	} else {
		tasks = store.ByStatus(task.Status(args[0]))
	}

	for _, t := range tasks {
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
	}
	return nil
}