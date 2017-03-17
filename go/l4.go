////////////////////////////////////////////////////////////////
// Filename: l4.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-01-10
// Author  : Atom
// Scope   :
// Usage   : constants & type conversion
////////////////////////////////////////////////////////////////

package main

import "fmt"
import "math"

func main() {

  const s = "Test"
  fmt.Println("S: ", s)

  var i = 20
  const j = 10

  // const k = j/i      // Not Gonna work, all var must be constants

  fmt.Println("Sin: ", math.Sin(float64(i)));
  fmt.Println("Sin: ", int(math.Sin(float64(i))));

}
