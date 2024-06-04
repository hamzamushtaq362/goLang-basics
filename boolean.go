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
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is lessthen 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is 45")
	}

	names := []string{"h1", "h2"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("cont", index)
			continue
		}
		if index > 2 {
			fmt.Println("cont", index)
			break
		}

		fmt.Println("the value", value)
	}
}
