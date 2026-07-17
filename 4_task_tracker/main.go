package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gaurav4288/4_task_tracker/commands"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"add": {
			name:        "add",
			description: "add a new task",
			callback:    commands.AddTask,
		},
		"update": {
			name:        "update",
			description: "update a task's description",
			callback:    commands.UpdateTask,
		},
		"delete": {
			name:        "delete",
			description: "delete a task",
			callback:    commands.DeleteTask,
		},
		"list": {
			name:        "list",
			description: "list tasks, optionally filtered by status",
			callback:    commands.ListTasks,
		},
		"mark-in-progress": {
			name:        "mark-in-progress",
			description: "mark a task as in progress",
			callback:    commands.MarkInProgress,
		},
		"mark-done": {
			name:        "mark-done",
			description: "mark a task as done",
			callback:    commands.MarkDone,
		},
		"exit": {
			name: 		 "exit",
			description: "end the cli",
			callback: 	 commands.Exit,
		},
	}
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("task-cli ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command:", commandName)
			continue
		}

		if err := command.callback(args...); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}