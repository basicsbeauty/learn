////////////////////////////////////////////////////////////////
// Filename: l6.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-01-10
// Author  : Atom
// Scope   :
// Usage   : if-else
////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
  i := 10
  if i % 2 == 0 {
    fmt.Println("Even in")
  } else {
    fmt.Println("Out In")
  }

  // Variables will available for the remaining scope
  if num:=9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "single digit")
  } else {
    fmt.Println(num, "double strength")
  }
}
