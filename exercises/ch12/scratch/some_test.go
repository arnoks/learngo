package scratch

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect1(t *testing.T) {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()

	fmt.Printf("x %d\n", x)
	fmt.Printf("a %v\n", a)
	fmt.Println(a.CanAddr())

	fmt.Printf("b %v\n", b)
	fmt.Println(b.CanAddr())

	fmt.Printf("d %v\n", d)
	fmt.Println(d.CanAddr())

}
