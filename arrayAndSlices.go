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
	// Array of numbers
	var ages [3]int = [3]int{1, 2, 3}
	// Or
	var ages2 = [3]int{1, 2, 3}
	fmt.Println(ages, len(ages))

	// Array of string
	names = [4]string{"t1", "t2", "t13", "t4"}
	// We can update array index
	names[1] = "update"

	// Slices (use array under the hood)
	var scores = []int{1, 2, 3}
	scores[2] = 4
	// we can also append
	scores = append(scores, 5)

	// Slice ranges

}
