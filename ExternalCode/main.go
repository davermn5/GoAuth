package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}

// To use, then run: go mod tidy (finds & downlaads the parent module [rsc.io/quote v1.5.2] (for this package that you imported below: import "rsc.io/quote" )
//	The corresponding modules' "go.mod" is saved in the system's global 'Module Cache'
//	Go computes a hash (from either the go.mod file itself or the modules' direct or indirect dependencies), and
//		then compares the computed hash against the corresponding hash value (which were calculated from the modules' direct or indirect dependencies) inside the "go.sum" file (for that module)
//  	This 'go.sum' file is used in authenticating the module.

// go list -m all  (The '-m' flag causes the 'list' command to list modules instead of packages
