package main

import (
	"log"
	"time"
)

func main() {
	defer trace("slowOperation")()

	time.Sleep(10 * time.Second)
}

func trace(name string) func() {
	start := time.Now()

	log.Printf("entering %s", name)

	return func() {
		log.Printf("exiting %s after %s seconds", name, time.Since(start))
	}
}
