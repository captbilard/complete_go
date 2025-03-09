package main

import (
	"fmt"
)

func main() {
	var (
		name       string = "Bilard"
		age        int    = 25
		isEmployed bool   = true
		city       string = "Calgary"
		country    string = "Canada"
	)

	fmt.Printf("My name is %s, I am %d years old. I live in %s which is located in %s\n", name, age, city, country)
	fmt.Printf("Is he employed? %t\n", isEmployed)

	// this is how you define enums in go, we can use the iota
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
	)
	fmt.Printf("These are the current months: Jan %d - Feb %d - Mar %d - Apr %d - May %d\n", Jan, Feb, Mar, Apr, May)

	subtraction := subtract(4, 7)
	fmt.Println("This is the subtraction result", subtraction)

	sum, _ := productAndSum(5, 12)
	fmt.Println("This is sum only", sum)

	_, product := productAndSum(5, 12)
	fmt.Println("This is product only", product)

	combinedSum, combinedProduct := productAndSum(5, 12)
	fmt.Printf("This is both sum %d and product %d\n", combinedSum, combinedProduct)
}

func subtract(a, b int) int {
	if a < b {
		fmt.Println("Param a must be greater than param b")
	}
	return a - b
}

func productAndSum(a, b int) (int, int) {
	return a + b, a * b
}
