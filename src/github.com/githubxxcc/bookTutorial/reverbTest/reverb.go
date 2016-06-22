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
		}

		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Timeouting")
			go handleConn(conn)
		}

	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)

	for input.Scan() {
		go eco(c, input.Text(), 1*time.Second)
	}

}

func eco(c net.Conn, shout string, d time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(d)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(d)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
