package countingwriter

import (
	"fmt"
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	gotw, gotc := CountingWriter(os.Stdout)
	fmt.Fprintln(gotw, "Hello World")
	fmt.Fprintln(gotw, "Hello World")
	fmt.Fprintln(gotw, "Hello World")
	fmt.Fprintln(gotw, "Hello World")
	fmt.Printf("Counted: %v", *gotc)
}
