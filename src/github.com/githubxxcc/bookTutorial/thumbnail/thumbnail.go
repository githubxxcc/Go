package main

import (
	"log"

	"gopl.io/ch8/thumbnail"

	"os"
	"sync"
)

func makeImage1(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Fatal(err)
		}
	}
}

func makeImage2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeImage3(filenames []string) {
	ch := make(chan int)

	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- 1
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func makeImage4(filenames []string) error {
	errorChan := make(chan error)

	for _, f := range filenames {
		go func(f string) { // no need to include chan error as a parameter, value capturing
			if _, err := thumbnail.ImageFile(f); err != nil {
				errorChan <- err
			}
		}(f)
	}

	for range filenames {
		if i <- errorChan; i != nil {
			return i // be careful of leaking
		}
	}

	return nil
}

func makeImage5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		file string
		err  error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var x item
			x.file, x.err = thumbnail.ImageFile(f)
			ch <- x
		}(f)
	}

	for range ch {
		x := <-ch

		if x.err != nil {
			return nil, x.err
		}

		thumbfiles = append(thumbfiles, x.file)
	}

	return thumbfiles, nil

}

func makeImage6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()

			file, err := thumbnail.ImageFile(f)

			if err != nil {
				return
			}

			infor, _ := os.Stat(file)

			sizes <- infor.Size()

		}(f)
	}

	//closer

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total

}
