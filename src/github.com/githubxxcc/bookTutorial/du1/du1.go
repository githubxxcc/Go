package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	sizes := make(chan int64)

	go func() {
		for _, root := range roots {
			walkDir(root, sizes)
		}
		close(sizes)
	}()

	var total, nfiles int64
	for size := range sizes {
		total += size
		nfiles++
	}

	fmt.Printf("%d files with size of %.1f \n", nfiles, float64(total)/1e9)

}

func walkDir(dir string, ch chan int64) {
	for _, file := range parseDir(dir) {
		if file.IsDir() {
			subdir := filepath.Join(dir, file.Name())
			walkDir(subdir, ch)
		} else {
			ch <- file.Size()
		}
	}
}

func parseDir(dir string) []os.FileInfo {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}

	return files
}
