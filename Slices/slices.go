package main

import "fmt"

func main() {

	//Here, we are just declaring (but not initializing) a slice (without using 'make()')
	// When we declare a slice, by default:
	//  - The slice does not contain any elements (length is zero), and
	//  - The slice is equal to NULL
	// OUTPUT: []  (If this slice contained 'string' elements, you'll notice the tiny gap representing a 'nil' size
	var s []int
	fmt.Println("uninit:", s, s == nil, len(s) == 0, len(s))

	// However, using 'make()' creates an empty slice with non-zero length, and each element's value is zero'ed.
	//  Here:
	//  - The slice contains 3 elements, and each element's value is set to 0.
	//  - The length is 3, and
	//  - The slice is NOT NULL
	// OUTPUT: [0]  (If this slice contained 'string' elements, you'll notice a larger gap e.g. OUTPUT: [   ] representing '3' empty strings)
	t := make([]int, 3)
	fmt.Println("t:", t, t == nil, len(t) == 0, len(t))

	fmt.Println("Does t[0] equal to zero:", t[0] == 0, "What is the value of t[0]:", t[0])
	//fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	/*
		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])
		fmt.Println("len:", len(s))

		s = append(s, "d")
		s = append(s, "e", "f")
		fmt.Println("apd:", s)

		c := make([]string, len(s))
		copy(c, s)
		fmt.Println("cpy:", c)

		l := s[2:5]
		fmt.Println("sl1:", l)

		l = s[:5]
		fmt.Println("sl2:", l)

		l = s[2:]
		fmt.Println("sl3:", l)

		t := []string{"g", "h", "i"}
		fmt.Println("dcl:", t)
	*/

	// You will notice that 'twoD' is declared & initialized, however the 3 sub-slices(e.g. sub-arrays) that are created within are UNinitialized - meaning that the value of each 'twoD[i]' is NIL.
	//  See step 2/3 below to find out how to initialize the sub-slice so that the values become zero'ed out (No longer NIL)
	twoD := make([][]int, 3)
	fmt.Println("twoD:", twoD, "twoD is nil?:", twoD == nil, "twoD[0] is nil?:", twoD[0] == nil, "twoD[0]:", twoD[0]) // OUTPUT: []  So, the zero'th index points to the 1st of the 3 empty sub-slices.

	// 1/3: So, if we wish to make "twoD[0]"'s inner 0 index set to Zero, then intuitively we would first attempt to do:
	//twoD[0][0] = 0 //This would work in PHP, but NOT in Go.
	//fmt.Println("twoD[0][0]:", twoD[0][0]) //Unfortunately, this creates a 'panic: runtime error: index out of range'

	// 2/3: Instead, we need to invoke the 'make()' method on the sub-slice, so that we CAN set the 'twoD[0][0]' element value to 0:
	twoD[0] = make([]int, 1) //twoD[0] represents the NIL sub-slice which was already created earlier, so now all is left is to just initialize the sub-slice (so that it receives a zeroe'd out value AND so that it's no longer NIL and can therefore hold actual non-NIL values (e.g. integers) that we can later set.
	//In other words, the output of the 'make()' method returns an initialized slice to our already declared 'twoD[0]', so that now 'twoD[0]' becomes declared and initialized.
	fmt.Println("twoD[0][0]:", twoD[0][0]) //Now outputs 0.

	// 3/3 Onwards with the rest of the lesson:
	anotherTwoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		anotherTwoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			anotherTwoD[i][j] = i + j
		}
	}
	fmt.Println("anotherTwoD:", anotherTwoD) // [ [0], [1,2], [2,3,4]  ]

	/*
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
		}
	*/
}
