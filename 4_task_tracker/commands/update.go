package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gaurav4288/4_task_tracker/internal/task"
)

func UpdateTask(args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: task-cli update <id> <new description>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid id: %s", args[0])
	}
	description := strings.Join(args[1:], " ")

	store, err := task.Load()
	if err != nil {
		return err
	}

	if !store.Update(id, description) {
		return fmt.Errorf("no task found with id %d", id)
	}

	if err := store.Save(); err != nil {
		return err
	}

	fmt.Printf("Task %d updated successfully\n", id)
	return nil
}
