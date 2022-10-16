package main

import (
	"fmt"
	"mgrep/dirsearch"
	"os"
	"sync"
)

//--Requirements:
//* Use goroutines to search through the files for a substring match
//* Display matches to the terminal as they are found
//  * Display the line number, file path, and complete line containing the match
//* Recurse into any subdirectories looking for matches
//* Use any synchronization method to ensure that all files
//  are searched, and all results are displayed before the program
//  terminates.
//
//--Notes:

func main() {

	inp := os.Args
	if len(inp) < 3 {
		fmt.Println("Please provide search arguments mgrep <file> <folder>")
		return
	}

	dir := inp[2]
	var fs chan string = make(chan string)
	var fwg sync.WaitGroup
	fwg.Add(1)
	go dirsearch.GetDirsAndFiles(dir, &fs, &fwg) // I cannot add to the channel if its not ready to be lisened from so havving goroutine is a must for code to work

	for {
		select {
		case file := <-fs:
			fmt.Println(file)
		}
		fwg.Wait()
	}
}
