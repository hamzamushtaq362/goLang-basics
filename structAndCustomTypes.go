// we want to run our program on our system
package main

/*
Package for formatting string
*/
import "fmt"

/*
entry point of application
If we rename this function, it wil not work
*/
func main() {
	myBill := newBill("Hamza Bill")
	fmt.Println("printing", myBill)
}
