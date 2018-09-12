package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "specify the listen port of the server")

func main() {
	flag.Parse()
	adress := fmt.Sprintf("localhost:%s", *port)
	fmt.Printf("Listening on adress: %s\n", adress)
	listener, err := net.Listen("tcp", adress)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(time.ANSIC)+"\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
