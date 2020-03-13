package scratch

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

func TestSprint(t *testing.T) {
	fmt.Println(Sprint("Eins"))
	fmt.Println(Sprint(true))
	fmt.Println(Sprint(3))
}

type myType struct {
	a string
	b int
	c float64
}

func TestReflekt(_ *testing.T) {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	fmt.Printf("%T\n", 3)
	v := reflect.ValueOf(3)
	rt := v.Type()
	fmt.Println(rt.String())

	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	var m myType
	fmt.Println(reflect.TypeOf(m))
}

func TestStructs(t *testing.T) {
	m := myType{"a", 42, math.Pi}
	fmt.Println(m)
	vom := reflect.ValueOf(m)
	fmt.Println(tom)
	fmt.Printf("Struct has %d Fields\n", vom.NumField())
	for i := 0; i < vom.NumField(); i++ {
		fmt.Println(vom.Field(i), vom.Type().Field(i).Name)
	}
}
