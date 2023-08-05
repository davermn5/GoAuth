package main

import "fmt"

func main() {

	//Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { //Notice that the '_' is present, which indicates the intent to use the value (but not the key)..
		sum += num
	}
	fmt.Println("sum:", sum)

	// i represents the index, starting at 0
	// num represents the value of the index
	//Notice that the '_' (blank identifier is absent); This signals an intent to use both key and value.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Reminds of being similar to foreach kvs as $k1 => $v1..
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs { // key, val
		fmt.Printf("%s -> %s\n", k, v) // OUTPUT: a -> apple \n  b -> banana
	}

	fruits := map[string]string{"c": "cherry", "d": "date"}
	for key := range fruits {
		fmt.Println("key:", key)
	}

	// Using 'range' on strings iterates over Unicode code points.
	//  The first value is the starting byte index of the 'rune' and
	//  the second the 'rune' itself
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}
