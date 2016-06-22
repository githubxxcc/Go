package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConn(conn)
	}
}

func eco(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	defer c.Close()

	for input.Scan() {
		// shout := input.Text()
		// delay := 1 * time.Second

		// fmt.Fprintln(c, "\t", strings.ToUpper(shout))
		// time.Sleep(delay)

		// fmt.Fprintln(c, "\t", shout)
		// time.Sleep(delay)

		// fmt.Fprintln(c, "\t", strings.ToLower(shout))

		// func(c net.Conn, shout string, delay time.Duration) {
		// 	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
		// 	time.Sleep(delay)

		// 	fmt.Fprintln(c, "\t", shout)
		// 	time.Sleep(delay)

		// 	fmt.Fprintln(c, "\t", strings.ToLower(shout))
		// }(c, input.Text(), 2*time.Second)
		text := input.Text()

		go eco(c, text, 2*time.Second)
	}
}
