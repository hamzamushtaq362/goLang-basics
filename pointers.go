// we want to run our program on our system
package main

/*
Package for formatting string
*/
import "fmt"

func updateByPointer(x *string) {
	*x = "Ali"
}

/*
entry point of application
If we rename this function, it wil not work
*/
func main() {
	name := "Hamza"

	m := &name
	// Memory Address
	fmt.Println("address", m)
	// Value at Pointer
	fmt.Println("Pointer is ", *m)

	// We are passing pointer
	updateByPointer(m)
	fmt.Println("So updated name is ", name)
}
