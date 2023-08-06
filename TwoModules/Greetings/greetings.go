package Greetings

import (
	"errors"
	"fmt"
)

// Capitalizing the 'H' in Hello allows this function to be invoked by other functions that are outside of this 'Greetings' package
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//Adding nil as the second value in the return, means that the function succeeded.
	return fmt.Sprintf("Hi, %v. Welcome!", name), nil
}
