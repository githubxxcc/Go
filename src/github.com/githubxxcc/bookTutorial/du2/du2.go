package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "printing out progress")

func main() {
	start := time.Now()
	//background function
	sizes := make(chan int64)
	flag.Parse()

	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		for _, dir := range roots {
			walkDir(dir, sizes)
		}
		close(sizes)
	}()

	tick := make(<-chan time.Time)

	if *verbose {
		tick = time.Tick(1000 * time.Millisecond)
	}

	var total, nfiles int64

loop:
	for {
		select {
		case size, ok := <-sizes:
			if !ok {
				break loop
			}

			nfiles++
			total += size
		case <-tick:
			printProgress(nfiles, total)
		}
	}

	printProgress(nfiles, total)
	fmt.Printf("using time %0.2f seconds\n", time.Since(start).Seconds())
}

func printProgress(nfiles, total int64) {
	fmt.Printf("Checked %d files, in total %.1f GB \n", nfiles, float64(total)/1e9)
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
