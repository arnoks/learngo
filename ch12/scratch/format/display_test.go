package format

import (
	"fmt"
	"math"
	"testing"
)

type theStruct struct {
	a int
	b float32
	c string
}

func TestAny(t *testing.T) {
	fmt.Println(Any(3))
	fmt.Println(Any(true))
	fmt.Println(Any(6.3))
	fmt.Println(Any([]string{"1", "2", "3"}))
	fmt.Println(Any([]int{1, 2, 3}))
	fmt.Println(Any(theStruct{1, math.Pi, "Hello World"}))

}

func TestDisplay(t *testing.T) {
	Display("theStruct", theStruct{1, math.Pi, "Hello World"})
}
