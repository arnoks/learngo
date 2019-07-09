// timing example from idomatic go

package main

import (
	"log"
	"time"
)

func main() {
	FunkyFunc()
}

// StartTimer is an example for timing functions
func StartTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Now().Sub(t)
		log.Println(name, "took", d)
	}
}

// FunkyFunc uses StartTimer to time a funnction call
func FunkyFunc() {
	stop := StartTimer("FunkyFunc")
	defer stop()
	time.Sleep(1 * time.Second)
}
