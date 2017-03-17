////////////////////////////////////////////////////////////////
// Filename: l2.go
// Location: /Users/Atom/repos/learn/go
// Project :
// Date    : 2017-03-10
// Author  : Atom
// Scope   :
// Usage   : Values & Expressions
////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {

	fmt.Println("Test" + "Message")

	fmt.Println("2+2:", 2+2)
	fmt.Println("7/3:", 7/3)
	fmt.Println("7/2.5:", 7/2.5)
	fmt.Println("7.5/2:", 7.5/2)
	fmt.Println("7.0/2.5:", 7.0/2.5)

	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(true || false)
	fmt.Println(true && false)

}
