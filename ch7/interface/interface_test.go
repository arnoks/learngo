package main_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestIng(t *testing.T) {
	t.Log("Runnning some arbitrary test!")
}

func TestIf1(t *testing.T) {
	fmt.Println("go Interface")
	var w io.Writer
	fmt.Printf("Initial: %v\n", w)
	w = os.Stdout
	fmt.Printf("Stdout: %v\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("byte.Buffer:%v\n", w)
	w = nil
	fmt.Printf("nil: %v\n", w)
}

func TestPanic(t *testing.T) {
	var w io.Writer
	w.Write([]byte("hello"))
}
