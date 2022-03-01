package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCelsiusFlag(t *testing.T) {
	cf := CelsiusFlag("temperature", 20.0, "Please provide temperature as xx[°]*[C|K|F]+")
	fmt.Printf("Temperatur: %s", cf)
}

func TestFlags(t *testing.T) {
	fmt.Println("Executing Tests ... ")
	fmt.Printf("Args: %v", os.Args[1:])
}
