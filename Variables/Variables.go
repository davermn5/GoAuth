package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//Variables that are declared without a corresponding
	// initialization are zer-valued. In this case, int defaults
	// to 0.
	var e int
	fmt.Println(e)

	//Shorthand for var f string = "apple"
	f := "apple"
	fmt.Println(f)
}
