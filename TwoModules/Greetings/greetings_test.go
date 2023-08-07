package Greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls Greetings.Hello with a name, checking
//
//	for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Fisher"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Fisher")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Fisher") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls Greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

//How to use: cd to the Greetings directory and then run: go test -v (which invokes the test)
// You can omit the -v flag, and instead the output will include results for only the tests that failed (esp. if you have a lot of tests).
