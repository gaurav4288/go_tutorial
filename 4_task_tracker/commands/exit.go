package commands

import (
	"fmt"
	"os"
)

func Exit(args ...string) error {
	fmt.Println("Closing the Task Tracker... Goodbye!")
	os.Exit(0)
	return nil
}
