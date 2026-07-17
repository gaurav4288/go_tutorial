package commands

import (
	"fmt"
	"strconv"

	"github.com/gaurav4288/4_task_tracker/internal/task"
)

func DeleteTask(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: task-cli delete <id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid id: %s", args[0])
	}

	store, err := task.Load()
	if err != nil {
		return err
	}

	if !store.Delete(id) {
		return fmt.Errorf("no task found with id %d", id)
	}

	if err := store.Save(); err != nil {
		return err
	}

	fmt.Printf("Task %d deleted successfully\n", id)
	return nil
}