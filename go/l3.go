////////////////////////////////////////////////////////////////
// Filename: l3.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-03-10
// Author  : Atom
// Scope   :
// Usage   : Variables
////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

  // Explicit
  var i string = "Test"
  fmt.Println("S: ", i)

  // Inference
  var c = "Message"
  fmt.Println("M: ", c)

  var d = true
  fmt.Println("D: ", d)

  // Implicit
  f := 12.3
  fmt.Println("F: ", f)

  // Deffaults
  var j int       // 0
  var b bool      // false
  var s string    // ""
  fmt.Println("F: ", j)
  fmt.Println("S: ", s)
  fmt.Println("B: ", b)
  fmt.Println("F: ", j)
  fmt.Println("S: ", s)
  fmt.Println("B: ", b)
}
