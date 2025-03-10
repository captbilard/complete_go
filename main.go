package main

import "fmt"

type name struct {
	shortName, longName string
}

func main() {
	objectArray := [5]name{}
	objectArray[2] = name{shortName: "boy", longName: "girl"}
	fmt.Printf("This is the object array %v\n", objectArray)

	sellableRates := map[string]string{
		"13": "RAC",
		"10": "LLK",
		"24": "PROMO",
	}
	rateType, exists := sellableRates["10"]
	fmt.Printf("does the rate type %s exist:%t\n", rateType, exists)
}
