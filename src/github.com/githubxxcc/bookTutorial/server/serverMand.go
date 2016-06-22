package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/githubxxcc/bookTutorial/gif"
	"github.com/githubxxcc/bookTutorial/mandelbrot"
)

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/cycle", handler2)
	http.HandleFunc("/mandelbrot", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {

	host := string(r.URL.Path)

	fmt.Fprintf(w, "USING URL %s", host)

	if index := strings.Index(host, "="); index != -1 {
		num, err := strconv.Atoi(host[index:])
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
		}

		gif.Lissajous(w, float64(num))
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	gif.Lissajous(w, float64(5))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	mandelbrot.GetMandelbrot(w)
}
