package main
import (
  "fmt"
)

func main()  {
  // go support 4 different variable declaration styles

  // 1. The short for for loop or small function
  s1 := ""
  fmt.Println(s1)

  // The Standard form with explicit type definition
  var s2 string
  fmt.Println(s2)

  // Similar to the first one but with explicit initialization and implicit
  // type definition
  var s3 = ""
  fmt.Println(s3)

  // Th most verbose defintion  with explicit type and initialization
  var s4 string = ""
  fmt.Println(s4)
}
