package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var clocks []string
	if len(os.Args) > 1 {
		clocks = os.Args[1:]
	} else {
		log.Fatal("No Adress")
	}
	for _, a := range clocks {
		conn, err := net.Dial("tcp", a)
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go mustCopy(os.Stdout, conn)
	}
	wg.Wait()
}

func mustCopy(dst io.Writer, src net.Conn) {
	defer src.Close()
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	wg.Done()
}
