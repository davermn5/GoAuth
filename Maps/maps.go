package main

import "fmt"

func main() {

	//Note: No length is being explictly defined for this 'map' type
	m := make(map[string]int)

	//Note also: When declaring & initializing (without using 'make()'):
	l := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", l)

	m["k1"] = 7
	m["k2"] = 13
	//Note: Getting the value of an unknown field e.g. m["k3"], returns 0 (see below)..

	fmt.Println("map:", m) //OUTPUT:  map: map[k1:7, k2:13]

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// Example of using the "blank identifier" (the underscore)
	//  //Prefixing the variable 'prs' returns TRUE | FALSE which indicates if the key was present in the map, and
	//  if the key was present then set 'prs' to TRUE, else FALSE.
	//  Note: The "blank identifier" is used to disambiguate between missing keys and keys with zero values like 0 or ""
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

}
