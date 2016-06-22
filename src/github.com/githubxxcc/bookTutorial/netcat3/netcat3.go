package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	done := make(chan struct{})

	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)

	conn.Close()

	// calling CloseWrite() allows the os.Stdout to keep printing all the ecos
	// c := conn.(*net.TCPConn)

	// c.CloseWrite()

	<-done

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
