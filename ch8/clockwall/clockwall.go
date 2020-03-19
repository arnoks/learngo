// tcp client reading the clocks from various
// clock servers and displays them in a table
// Execercise 8.1
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
		log.Fatal("No Address")
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
