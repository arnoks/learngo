package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	fmt.Println("Commencing Countdown ...")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		<-tick
	}
	fmt.Println("Lauching Rocket!")
}
