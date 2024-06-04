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
	menu := map[string]float64{
		"pizza":  22.33,
		"burger": 12.99,
	}
	// Print
	fmt.Println(menu["pizza"])

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	phonebook := map[int]string{
		1010: "Hamza",
		2010: "Ahmed",
	}

	for key, value := range phonebook {
		fmt.Println(key, "-", value)
	}

	// update the item
	phonebook[1010] = "Ali"
	fmt.Println(phonebook[1010])
}
