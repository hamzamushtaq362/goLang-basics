// we want to run our program on our system
package main

/*
Package for formatting string
*/
import "fmt"

func update(x string) {
	x = "Ali"
}

func forceUpdate(x string) string {
	x = "Ali"
	return x
}

/*
entry point of application
If we rename this function, it wil not work
*/
func main() {
	name := "Hamza"

	// simple update
	update(name)

	// Print
	fmt.Println("my name is", name)

	// update by returning value
	name = forceUpdate(name)
	fmt.Println("my name is", name)
}
