package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//If you only want a subset of the returned values, use the blank identifier '_'
	_, c := vals() //Grabs the 7 from the function, based on it's argumental placeholder location (2nd spot from left)

	//You could also do: (which grabs the 3 instead)
	//c, _ := vals()
	fmt.Println(c)
	//Interesting output: %!s(int=3) and %!s(int=7)
	//fmt.Printf("%s and %s\n", a, b)
}
