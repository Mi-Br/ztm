//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type LetterCount struct {
	number int
	sync.Mutex
}

func readInput() []string {
	fmt.Println("Provide text and hit enter to count the words")
	rd := bufio.NewReader(os.Stdin)
	txt, err := rd.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return []string{}
	} else {
		return strings.Split(txt, " ")
	}
}

func countLetters(wg *sync.WaitGroup, lc *LetterCount, word string) {
	count := len(word)
	lc.Lock()
	defer wg.Done()
	defer lc.Unlock()
	lc.number += count
	fmt.Printf("Routine finished for, %v. Counted + %v \n ", word, count)
}

func main() {
	var lcount LetterCount
	var wg sync.WaitGroup
	words := readInput()

	for i := 0; i < len(words); i++ {
		wg.Add(1)
		go countLetters(&wg, &lcount, words[i])
	}
	wg.Wait()
	fmt.Println("\n\nRoutines finished, final count: ", lcount)
}
