package dirsearch

import (
	"fmt"
	"os"
	"path"
	"sync"
)

// GetDirsAndFiles lists content of a given path
// It updates stack of Directories and Files.
func GetDirsAndFiles(p string, fch *chan string, wg *sync.WaitGroup) {
	files, error := os.ReadDir(p)
	wg.Done()
	if error == nil {
		for _, f := range files {
			if f.IsDir() {
				fmt.Println("added dir")
				// *dch <- fmt.Sprint(path.Join(p, f.Name()))
				wg.Add(1)
				go GetDirsAndFiles(path.Join(p, f.Name()), fch, wg)
			} else {
				fmt.Println("added file")
				*fch <- fmt.Sprint(path.Join(p, f.Name()))
			}
		}
	} else {
		fmt.Println("error:", error)
	}
}
