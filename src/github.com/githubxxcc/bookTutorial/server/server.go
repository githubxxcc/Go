package main

import(
"net/http"
"log"
"github.com/githubxxcc/bookTutorial/gif"
"strconv"
"fmt"
"strings"

)

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/cycle", handler2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request){
	
	host := string(r.URL.Path)

	fmt.Fprintf(w, "USING URL %s", host)

	if index := strings.Index(host, "="); index != -1{
		num, err := strconv.Atoi(host[index:])
		if err != nil{
			fmt.Fprintf(w, "Error %s", err)
		}

		gif.Lissajous(w, float64(num))
	}
}

func handler1(w http.ResponseWriter, r *http.Request){
	gif.Lissajous(w, float64(5))
}