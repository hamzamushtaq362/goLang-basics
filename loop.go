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
	// While Loop
	x := 0
	for x < 5 {
		fmt.Println("value", x)
		x++
	}

	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("value", i)
	}

	//
	names := []string{"h2", "h1", "h3"}
	for i := 0; i < 5; i++ {
		fmt.Println("name is", names[i])
	}

	// For in Loop
	for index, value := range names {
		fmt.Printf("position is %v and value is %v", index, value)
	}

	// For in Loop
	for _, value := range names {
		fmt.Printf("value is %v \n", value)
	}
}
