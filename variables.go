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
	// Print
	fmt.Println("Hello, World!")

	// strings
	var nameOne string = "test1"
	// we need to use it otherwise error
	fmt.Println(nameOne)

	// We can also use it without giving type
	var nameTwo = "test2"
	fmt.Println((nameTwo))

	// We can declare varibalewith type, but without value
	var nameThree string
	fmt.Println((nameThree))

	// Re-Declare value
	nameThree = "test3"
	fmt.Println(nameThree)

	// Integer
	var ageOne int = 30
	var ageTwo = 30
	var ageThree int
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	// Specify range of numbers allowed
	var numOne: int8 = 25
	// Specify range of numbers allowed
	var numTwo: int16 = 25
	// Means we cannot have negative number
	var numThree uint = 28

	// Float
	var scoreOne float32=-1.5
	var scoreTwo float64=823455253452537.7
	scoreThree:=1.5
}
