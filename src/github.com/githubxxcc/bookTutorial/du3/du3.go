package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "printing progress")

func main() {
	sizes := make(chan int64)

	start := time.Now()

	var wg sync.WaitGroup

	flag.Parse()
	roots := flag.Args()
	tick := make(<-chan time.Time)

	if len(roots) == 0 {
		roots = []string{"."}
	}

	for _, dir := range roots {
		wg.Add(1)

		go walkDir(dir, sizes, &wg)
	}

	if *verbose {
		tick = time.Tick(1000 * time.Millisecond)
	}

	var nfiles, ndata int64
	//closer

	go func() {
		wg.Wait()
		close(sizes)
	}()

loop:
	for {
		select {
		case <-tick:
			printProgress(nfiles, ndata)
		case value, ok := <-sizes:
			if !ok {
				break loop
			}
			nfiles++
			ndata += value
		}

	}

	printProgress(nfiles, ndata)
	fmt.Printf("using time %0.2f seconds\n", time.Since(start).Seconds())

}

func printProgress(nfiles, total int64) {
	fmt.Printf("Checked %d files, in total %.1f GB \n", nfiles, float64(total)/1e9)
}

func walkDir(dir string, ch chan int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, file := range parseDir(dir) {
		if file.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, file.Name())
			go walkDir(subdir, ch, wg)

		} else {
			ch <- file.Size()
		}
	}
}

var sema = make(chan struct{}, 50)

func parseDir(dir string) []os.FileInfo {
	sema <- struct{}{}

	defer func() { <-sema }()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}

	return files
}
