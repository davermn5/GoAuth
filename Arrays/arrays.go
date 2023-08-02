package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//Tip for declaring arrays: Reading from left to right..builds the array starting from the outside -> inside
	// When you declare a nested array, there will always be a parent array encapsulating the whole thing.
	//There are 4 sets of arrays with 3 integers each.
	//var multiArr [4][3]int
	//fmt.Println("multiArr:", multiArr)

	//There are 2 outside arrays.
	//  Each of those 2 outside arrays contain 4 arrays each having 3 integers.
	//var multiArr [2][4][3]int
	//fmt.Println("multiArr:", multiArr)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	// [ [0,1,2], [1,2,3] ]
	fmt.Println("2d: ", twoD)
}
