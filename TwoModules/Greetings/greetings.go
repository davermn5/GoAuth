package Greetings

import "fmt"

// Capitalizing the 'H' in Hello allows this function to be invoked by other functions that are outside of this 'Greetings' package
func Hello(name string) string {
	return fmt.Sprintf("Hi, %v. Welcome!", name)
}
