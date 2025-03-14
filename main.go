package main

import (
	"fmt"
)

type name struct {
	shortName, longName string
}

func main() {
	objectArray := [5]name{}
	objectArray[2] = name{shortName: "boy", longName: "girl"}
	fmt.Printf("This is the struct array with just the values %v\n", objectArray)
	fmt.Printf("This is the struct array with values & keys %+v\n", objectArray)

	newSlice := make([]int, 3, 5)
	newSlice[0] = 5
	fmt.Printf("This is our new slice %v\n", newSlice)

	sellableRates := map[string]string{
		"13": "RAC",
		"10": "LLK",
		"24": "PROMO",
	}
	rateType, exists := sellableRates["10"]
	fmt.Printf("does the rate type %s exist:%t\n", rateType, exists)

	// creating a struct and initialiazing immediately
	employee := struct {
		name       string
		salary     int
		department string
	}{
		name:       "Adebayo salami",
		salary:     65000,
		department: "Sunridge mall",
	}
	fmt.Printf("this is the employee details: %+v", employee)
}
