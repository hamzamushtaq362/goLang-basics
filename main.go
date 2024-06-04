// we want to run our program on our system
package main

/*
Package for formatting string
*/
import "fmt"

// avaliable globally
var scope = 123

/*
entry point of application
If we rename this function, it wil not work
*/
func main() {
	// avaliable locally
	var local = 12
	// Print
	fmt.Println("Hello, World!")

	// for this you need to yun both mian and scope filein terminal
	sayGreety("Hamza")

	// Bill object
	myBill := newBill("Hamza")
	// function associated with bill object
	fmt.Printf(myBill.format())
}
