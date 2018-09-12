package main

import (
	"fmt"
	"time"
)

func main() {
	wad_e_ma := time.Tick(10 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("Wad e ma en Aacheblegg")
		<-wad_e_ma
	}
}
