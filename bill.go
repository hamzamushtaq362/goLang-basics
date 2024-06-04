package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// we receive bill type in this function
func (b bill) format() string {
	fs := "Bill: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...pounds%v \n", k+":", v)
		total += v
	}

	// Total
	fs += fmt.Sprintf("%v ...Pounds", "total:", total)
	return fs
}
