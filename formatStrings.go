// we want to run our program on our system
package main

/*
formatting string
*/
import "fmt"

/*
entry point of application
If we rename this function, But it wil not work
*/
func main() {
	age := 25
	name := "John"

	// Print without any line
	fmt.Print("Hello ")
	// we need to add seprate lines
	fmt.Print(", World! \n")
	fmt.Print("new line \n")

	// We can use, to auto maticaly add line
	fmt.Println("Test 1")
	fmt.Println("Test 2")

	// use variable with strings
	fmt.Println("age is", age, "and name is ", name)

	// use variable with strings shortcut method
	// Printf
	fmt.Printf("my name is %v and age is %v \n", age, name)

	// add quotes on string
	fmt.Printf("my name is %q and age is %q \n", age, name)

	// get typeof variables
	fmt.Printf("my name is %T and age is %T \n", age, name)

	// Float
	fmt.Printf("my name is %f  \n", 2.55)

	// Float roundOff
	fmt.Printf("my name is %0.1f  \n", 2.55)

	// Save formatted strings
	var str = fmt.Sprintf("age is %v and name is %v \n", age, name)
	fmt.Println("New string is", str)
}
