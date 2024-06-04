// we want to run our program on our system
package main

import "fmt"

/*
Package for formatting string
*/

//
func sayGreetings(n string) {
	fmt.Printf("good mroning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("good Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(a float64, b float64) float64 {
	return a * b
}

/*
entry point of application
If we rename this function, it wil not work
*/
func main() {
	names := []string{"h1", "h2"}
	sayGreetings("John")
	sayGreetings("Alex")

	sayBye("Alex")
	cycleNames(names, sayGreetings)

	res := circleArea(2.33, 3.77)
	fmt.Printf("value is %v", res)
}
