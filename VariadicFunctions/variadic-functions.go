package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Within this function, the type of 'nums' is equivalent to '[]int' (slice of integers).
//
//	Because of that, we can't pass in array, only create a slice of integers as the named parameter to the receiving function argument.
func sum(nums ...int) { // The '...' is space-separated from the named argument such like: nums ...
	fmt.Print(nums, " ") //OUTPUT: [1 2]  (when invoking the function signature: sum(1, 2)
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

/*
// Note: '...' only supports slices This here says accept []map[string]string (top-level slice, not a top-level map)
func spaceIt(spacedMap ...map[string]string) {
	//spacedStr := ""
	for k, v := range spacedMap {
		fmt.Printf("%s -> %s\n", k, v)

	}
}
*/

/*
Function: intToString()
Accepts an arbitrary sequence of integers and returns a stringified version of that.

nums3 ...int   	[]int   A slice of integers
@return 		string	Returns a stringified version of the input integer(s)
*/
func intToString(nums3 ...int) string {
	strSlice := make([]string, len(nums3))
	for i, v := range nums3 {
		strSlice[i] = strconv.Itoa(v)
	}
	return strings.Join(strSlice, ",")
}

func main() {

	//Invoking each parameter manually...
	sum(1, 2)
	sum(1, 2, 3)
	sum(4, 5)

	//OR.. declaring and initializing a variable, and then passing the '<var>...' named parameter into the function signature.
	nums := []int{1, 2, 3, 4}
	sum(nums...) //Note: The '...' are right up against the 'nums...' as a named parameter.

	nums2 := []int{1, 2, 3, 4, 5}
	sum(nums2...)

	//This is how you create a new type (e.g. int and string are types). (This new type is not a variable)
	//type spacedMap map[string]string

	//spacedMapSlice := make([]map[string]string, 1, 1)
	//spaceIt(spacedMapSlice...)

	nums3 := []int{6, 7, 8, 9, 10}
	fmt.Println(intToString(nums3...))

}
