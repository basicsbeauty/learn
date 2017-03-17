////////////////////////////////////////////////////////////////
// Filename: l5.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-01-10
// Author  : Atom
// Scope   :
// Usage   : For
////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

  i := 5
  for i < 20 {
    fmt.Println("I: ", i)
    i++
  }

  // continue
  for j := 0;j<10;j++ {
    if j % 2 ==0 {
      continue
    }
    fmt.Println("J: ", j)
  }

  // break
  for {
    fmt.Println("Forever")
    break
  }
}
