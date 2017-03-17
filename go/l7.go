////////////////////////////////////////////////////////////////
// Filename: l7.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-03-10
// Author  : Atom
// Scope   :
// Usage   : switch
////////////////////////////////////////////////////////////////

package main

import "fmt"
import "time"

func main() {

  i := 3
  switch i {
    case 1: fmt.Println("It's One")
    case 2: fmt.Println("It's Two")
    case 3: fmt.Println("It's Three")
            fmt.Println("It's Three1")
  }

  fmt.Println("T: ", time.Now())
  fmt.Println("D: ", time.Now().Second())
  fmt.Println("D: ", time.Now().Day())
  fmt.Println("D: ", time.Now().Day())
  fmt.Println("M: ", time.Now().Month())
  fmt.Println("Y: ", time.Now().Year())
  fmt.Println("W: ", time.Now().Weekday())
  fmt.Println("W: ", time.Now().Weekday()+1)

  // Grouping
  switch time.Now().Weekday() {
    case time.Monday,
         time.Tuesday,
         time.Wednesday,
         time.Thursday,
         time.Friday:
           fmt.Println("D: Weekday")
    case time.Saturday,
         time.Sunday:
           fmt.Println("D: Weekend")
  }

  // Default
  // A switch without expression is same a if/else
  t := time.Now()
  switch  {
    case t.Second() % 2 == 0: fmt.Println("Even Second")
    default:
      fmt.Println("Odd Second")
  }

}
