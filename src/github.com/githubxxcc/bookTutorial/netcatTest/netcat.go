package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Done")
	}()

	mustCopy(conn, os.Stdin)

	c := conn.(*net.TCPConn)

	c.CloseWrite()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
