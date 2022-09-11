//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	// "io"
	"os"
	// "time"
)

type FileOpenError struct {
	msg, context string
}

func (e *FileOpenError) Error() string {
	return fmt.Sprint("Error opening file: %v, context %v", e.msg, e.context)
}

func readFileAndSum(name string, s *int) (err error) {
	line := 0
	file, err := os.Open(name)
	if err != nil {
		return &FileOpenError{msg: name, context: "@Opening file"}
	} else {
		r := bufio.NewReader(file)
		for {
			str, line_err := r.ReadString('\n')
			line++
			if line_err == nil {
				val, err := strconv.Atoi(strings.Trim(str, "\n"))
				if err == nil {
					*s += val
				} else {
					// fmt.Println(err)
					// fmt.Printf("Cannot convert value in file %v, line %v, skiping", name, line)
					continue
				}
			}
			if line_err == io.EOF {
				break
			}
		}
	}
	return nil
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}
	totalSum := 0
	fmt.Println("Syncing....")
	for _, f := range files {
		go readFileAndSum(f, &totalSum)
	}

	fmt.Println("Syncing....")
	time.Sleep(5 * time.Second)
	fmt.Println(totalSum)
}
