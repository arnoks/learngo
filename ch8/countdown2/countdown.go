package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing Countdown, Press Return to abort")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Aborting...")
			os.Exit(1)
		}
	}
	fmt.Println("Launching Rocket!")
}
