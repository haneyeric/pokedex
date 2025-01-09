package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scn := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scn.Scan() {
			input := scn.Text()
			word := cleanInput(input)[0]
			cmd, ok := commands[word]
			if !ok {
				fmt.Printf("%s is not a known command, try help for more info", word)
			} else {
				err := cmd.callback()
				if err != nil {
					fmt.Printf("Couldn't execute command: %s\n", err)
				}
			}
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}
