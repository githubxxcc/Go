package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	ch   chan string
	name string
}

var (
	entering  = make(chan client)
	leaving   = make(chan client)
	messaging = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	go broadcast()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func broadcast() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messaging:
			for client := range clients {
				client.ch <- msg
			}
		case cl := <-entering:

			clients[cl] = true
			printNames(clients)
		case cl := <-leaving:
			delete(clients, cl)
			close(cl.ch)
			printNames(clients)
		}
	}

}

func printNames(clients map[client]bool) {
	for cli := range clients {
		cli.ch <- fmt.Sprintln("People in the chat room:")
		for cliName := range clients {
			cli.ch <- fmt.Sprintln(cliName.name)
		}
	}
}

func handleConn(c net.Conn) {
	ch := client{ch: make(chan string), name: ""}
	go clientWriter(ch.ch, c)

	who := c.RemoteAddr().String()
	ch.name = who

	ch.ch <- "You are" + who
	messaging <- who + "has entered"
	entering <- ch

	var lastSent = time.Now()

	//disconnect client
	go func() {
		for {
			if time.Since(lastSent).Seconds() >= float64(10) {
				left(ch)
				c.Close()
			}
		}
	}()

	input := bufio.NewScanner(c)
	for input.Scan() {
		messaging <- who + ":" + input.Text()
		lastSent = time.Now()
	}

	left(ch)
	c.Close()

}

func left(cl client) {
	leaving <- cl.ch
	messaging <- cl.name + "is leaving"
}

func clientWriter(ch chan string, c net.Conn) {
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}
