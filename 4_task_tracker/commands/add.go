package commands

import (
	"fmt"
	"strings"

	"github.com/gaurav4288/4_task_tracker/internal/task"
)

func AddTask(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: task-cli add <description>")
	}
	description := strings.Join(args, " ")

	store, err := task.Load()
	if err != nil {
		return err
	}

	nextID := 1
	for _, t := range store.Tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	store.Add(task.Task{
		ID:          nextID,
		Description: description,
		Status:      task.StatusTodo,
	})

	if err := store.Save(); err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", nextID)
	return nil
}
