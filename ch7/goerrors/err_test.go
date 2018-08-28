package err_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	e1 := errors.New("Fehler")
	fmt.Printf("%v\n", e1)
	e2 := fmt.Errorf("Fehler")
	if e1 == e2 {
		fmt.Println("Fehler")
	} else {
		fmt.Println("different")
	}
}
