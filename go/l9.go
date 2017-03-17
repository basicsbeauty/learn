////////////////////////////////////////////////////////////////
// Filename: l9.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-03-11
// Author  : Atom
// Scope   : Slices
// Usage   :
////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

  s := make([]string, 3)
  fmt.Println("S: ", s)

  s[0] = "Test"
  s[1] = "*"
  s[2] = "Message"
  fmt.Println("S: ", s)
}
