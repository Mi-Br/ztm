//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	dict := []string{
		"Not feeling well today, you?",
		"How is weather today?",
		"Are you single?",
		"Hope things are going well for you?",
		"Are you alone?",
		"Would you like to chat or what?",
	}

	var (
		commandCount, linesCount int
	)

	rand.Seed(5)
	random := rand.Intn(len(dict))
	commandCount, linesCount = 0, 0
	for {
		fmt.Println("type something pls...")
		inp := os.Stdin
		r := bufio.NewReader(inp)

		input, inpErr := r.ReadString('\n')
		if inpErr == nil {
			commandCount += 1
			if len(input) > 0 && input[0] != byte('\n') {
				linesCount += 1
			}
			switch strings.TrimSpace(input) {
			case "q":
				fmt.Printf("Closing applicaiton, number of new lines: %v, number of commands: %v", linesCount, commandCount)
				return
			case "Q":
				fmt.Printf("Closing applicaiton, number of new lines: %v, number of commands: %v", linesCount, commandCount)
				return
			case "hello":
				fmt.Println(dict[random])
			case "bye":
				fmt.Println(dict[random])
			default:
				continue
			}
		} else {
			fmt.Println("Something went wrong, please try again (type <Q> or <q> followed by Enter - to quit")
			continue
		}
	}

}
