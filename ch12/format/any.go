//
package format

import (
	"reflect"
	"strconv"
)

// Any retuens a formated string
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

//
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return ("invalid")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	default:
		return v.Type().String() + " value not formatted"
	}
}
