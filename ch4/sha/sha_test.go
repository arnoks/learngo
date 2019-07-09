package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	main()
	t.Error("Forced fail") // force display of output while running the test
}
