package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var cleantext []string
	// Split the input text into words using strings.Fields
	// This will automatically trim leading and trailing whitespace
	// and split the text on any amount of whitespace
	cleantext = strings.Fields(text)
	// Convert each word to lowercase using strings.ToLower
	for i, word := range cleantext {
		cleantext[i] = strings.ToLower(word)
	}

	return cleantext
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}