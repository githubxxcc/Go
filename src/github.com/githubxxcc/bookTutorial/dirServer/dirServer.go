package main

import (
	"log"
	"net"
	"os"
)

func main() {
	server := os.Args[1:]

	tcpAdr, err := net.ResolveTCPAddr("tcp", server)
	checkErr(err)

	listener, err := net.Listen("tcp", tcpAdr)
	checkErr(err)

	for {
		conn, err := listener.Accept()

		checkErr(err)

		handleConn(conn)
	}

}

func checkErr(err error) {
	log.Fatal(err)
}

func handleConn(c net.Conn) {
	defer c.Close()

	var buf [512]byte

	for {
		n, err := c.Read(buf[0:])
		if err != nil {
			c.Close()
			return
		}
	}

}
