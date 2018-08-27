package main

import (
	"os"
	"testing"
)

// test to Sleep

func TestSleep(t *testing.T) {
	os.Args = []string{"sleep", "-period", "3sec"}
	main()
}
