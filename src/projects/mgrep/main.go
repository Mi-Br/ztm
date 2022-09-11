package main

import (
	"fmt"
	"io/fs"
	"os"
	"mgrep/filesearch"
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

type DirectoryStack []fs.DirEntry
type FileStack []fs.File

func ProcessDirectory(path string, Dirs *DirectoryStack, Files *FileStack) {
	files, error := os.ReadDir(path)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(files)
	}
}

func main() {

	dStack := DirectoryStack{}
	fStack := FileStack{}

	inp := os.Args
	dir := inp[2]
	fileSearch("whatever")
	list, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, f := range list {
			fmt.Println(f.Name(), f.IsDir())
		}
	}
	// list down context of directory

	// mgrep search_string search_dir
}
