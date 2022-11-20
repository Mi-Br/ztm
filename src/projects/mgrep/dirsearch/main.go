package dirsearch

import (
	"fmt"
	"os"
	"path"
)

// GetDirsAndFiles lists content of a given path
// It updates stack of Directories and Files.
func GetDirsAndFiles(p string, fch chan<- string, dch chan<- string) {
	files, error := os.ReadDir(p)
	if error == nil {
		for _, f := range files {
			if f.IsDir() {
				fmt.Println("added dir")
				dch <- fmt.Sprint(path.Join(p, f.Name()))
			} else {
				fmt.Println("added file")
				fch <- fmt.Sprint(path.Join(p, f.Name()))
			}
		}
	} else {
		fmt.Println("error:", error)
	}
}
