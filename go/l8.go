////////////////////////////////////////////////////////////////
// Filename: l8.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-03-11
// Author  : Atom
// Scope   : Array
// Usage   :
////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

  // Simple Array - Size & Type & Defaul Value
  var a[5] int
  fmt.Println("A: ", a)

  // Access & Assignment
  fmt.Println("A: ", a[4])
  a[4] = 783
  fmt.Println("A: ", a[4])
  fmt.Println("A: ", a)

  // Length
  fmt.Println("L: ", len(a))

  // Declaration and Initlization
  b := [5] int { 7, 8, 9, 10, 11}
  fmt.Println("B: ", b)

  // Two Dimensional Array
  var c[5][4] int
  fmt.Println("2d: ", c)

  for i:=0; i<5; i++ {
    for j:=0; j<4; j++ {
      c[i][j] = i+1*j+1
    }
  }

  fmt.Println("2d: ", c)
