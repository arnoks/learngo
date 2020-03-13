package format

import (
	"fmt"
	"testing"
)

func TestAny(t *testing.T) {
	fmt.Println(Any(3))
	fmt.Println(Any(true))
	fmt.Println(Any(6.3))
}
