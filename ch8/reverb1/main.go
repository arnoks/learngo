package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// reflect incoming data

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\n", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\n", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\n", strings.ToLower(shout))
	time.Sleep(delay)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	conn, err := net.Dial("")
}
