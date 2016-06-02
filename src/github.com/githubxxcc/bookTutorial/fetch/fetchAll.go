package main

import(
"fmt"
"os"
"io"
"io/ioutil"
"net/http"
"time"
)


func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}


	for range os.Args[1:]{
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


func fetch(url string, ch chan string){
	start := time.Now()

	req, err := http.Get(url)

	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, req.Body)

	req.Body.Close()

	if err != nil{
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
	}

	timePassed := time.Since(start).Seconds()

	ch <- fmt.Sprintf("From: %s, Bytes: %d, Time Used %.2f seconds", url, nbytes, timePassed )
}