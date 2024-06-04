// we want to run our program on our system
package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Package for formatting string
*/

/*
entry point of application
If we rename this function, it wil not work
*/
func main() {
	greeting := "Hello my name is"

	// strings is a package
	// Contains Method
	fmt.Println(strings.Contains(greeting, "Hello"))
	// ReplaceAll Method
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// ToUpper Method
	fmt.Println(strings.ToUpper(greeting))
	// Index Method (find position of letter)
	fmt.Println(strings.Index(greeting, "ll"))
	// Split Method (find position of letter and split)
	fmt.Println(strings.Split(greeting, " "))

	// sort is a Package
	ages := []int{1, 2, 3, 4, 5, 6}
	sort.Ints(ages)

}
