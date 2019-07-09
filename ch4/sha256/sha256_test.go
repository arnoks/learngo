package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Printf("executing %s\n", t.Name())
	main()
	t.Error("forced fail") // always fail to force output to be displayed
}
