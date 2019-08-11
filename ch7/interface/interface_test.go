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

func TestNilPointerWrite(t *testing.T) {
	var w io.Writer
	var bb bytes.Buffer
	w = bytes.NewBuffer(make([]byte, 100))
	n, err := w.Write([]byte("hello"))
	if n > 0 && err == nil {
		fmt.Printf("%v bytes written to buffer error:%v", n, err)
	}
	w = &bb
	if w != nil {
		n, err = w.Write([]byte("hello"))
		if n > 0 && err == nil {
			fmt.Printf("%v bytes written to buffer error:%v", n, err)
		}
	}
}

func Test_Error(t *testing.T) {
	var err error
	err = io.EOF
	fmt.Println(err)
	fmt.Println(io.ErrClosedPipe)
	fmt.Println(io.ErrNoProgress)
	fmt.Println(io.ErrUnexpectedEOF)
	fmt.Println(os.ErrClosed)
}
