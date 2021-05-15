package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int64, 50)
	squares := make(chan int64, 50)
	summer := make(chan int64, 50)
	startt := time.Now()
	go func() {
		var i int64
		for i = 0; i < 10000000; i++ {
			naturals <- i
		}
		close(naturals)
	}()
	go func() {
		for n := range naturals {
			squares <- n * n
		}
		close(squares)
	}()
	go func() {
		var sum int64 = 0
		for sq := range squares {
			sum += sq
			summer <- sum
		}
		close(summer)
	}()
	var i int64 = 0
	for s := range summer {
		i += s
	}
	fmt.Printf("Count: %d, duration: %v\n", i, time.Since(startt))
}
