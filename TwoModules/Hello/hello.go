// This file is the caller, which invokes the Hello function that has already been defined in the 'Greetings' module
package main

import (
	"Greetings" //Placing this package name here, and then running 'go mod tidy' (see below) creates the required dependency (to the Greetings module package code)
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names:
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

//Then run: go mod edit -replace Greetings=../Greetings
// And you will see that the 'hello.go' file was modified by this above command.

// Also, then run (from within the Hello directory): go mod tidy (which synchronizes the Hello module's dependencies, adding those required by the code, but not yet tracked in the module.
//  And you will see the hello.mod file now contains a 'require' directive with associated value.

// To compile, either: go build
// To compile and install on system:
//  go list -f '{{.Target}}' (copy the resultant "install" path) up to 'bin' e.g. /Users/<name>/go/bin
//   Then add this install path to your $PATH by editing the following file (sudo nano /etc/paths.d/newfile ...then drop the install path (from above) into this file called 'newfile'.
