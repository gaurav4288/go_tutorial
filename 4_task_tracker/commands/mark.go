package commands

import (
	"fmt"
	"strconv"

	"github.com/gaurav4288/4_task_tracker/internal/task"
)

func MarkInProgress(args ...string) error {
	return setStatus(args, task.StatusInProgress)
}

func MarkDone(args ...string) error {
	return setStatus(args, task.StatusDone)
}

func setStatus(args []string, status task.Status) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: task-cli mark-%s <id>", status)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid id: %s", args[0])
	}

	store, err := task.Load()
	if err != nil {
		return err
	}

	if !store.SetStatus(id, status) {
		return fmt.Errorf("no task found with id %d", id)
	}

	if err := store.Save(); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as %s\n", id, status)
	return nil
}