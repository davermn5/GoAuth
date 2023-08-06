package Greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Capitalizing the 'H' in Hello allows this function to be invoked by other functions that are outside of this 'Greetings' package
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//Adding nil as the second value in the return, means that the function succeeded.
	//Create a message using a random format:
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
