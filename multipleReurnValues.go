// we want to run our program on our system
package main

import (
	"fmt"
	"strings"
)

/*
Package for formatting string
*/

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, value := range names {
		// value[:1] get first letter of each string
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

/*
entry point of application
If we rename this function, it wil not work
*/
func main() {
	fn1, fn2 := getInitials("John Locket")
	fmt.Println(fn1, fn2)
}
