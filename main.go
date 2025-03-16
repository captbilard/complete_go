package main

import (
	"fmt"
)

type Person struct {
	shortName, longName string
}

// To add a method to the person struct, we define a func with parenthesis containing the struct type. Like so
func (p *Person) modifyPersonLongName(name string) {
	p.longName = name
	fmt.Printf("This is the new long name modified by the person struct method %v\n", p.longName)
}

func main() {
	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		name    string
		address Address
		phone   string
	}

	contact := Contact{
		name: "Merkley",
		address: Address{
			City:   "Lagos",
			Street: "Mushin",
		},
	}
	fmt.Printf("This is the contact details %v\n", contact)

	person := Person{
		shortName: "derek",
		longName:  "derek-Chukwu",
	}
	// here we modify properties of the person struct using both an outside defined function and
	// a struct method.
	fmt.Println("This is the old short name", person.shortName)
	modifyPersonShortName(&person)
	fmt.Println("This is the new short name after function execution", person.shortName)

	fmt.Println("This is the modification of person using the person method")
	person.modifyPersonLongName("Nkedelim")

	x := 20
	ptr := &x
	fmt.Printf("This is the value of x %d and this is the value of pointer gotten from x %p\n", x, ptr)
	*ptr = 30
	x = 50
	fmt.Printf("This is the new value of x %d and this is the value of pointer gotten from x %p\n", x, ptr)

}

// if we only pass in the person variable here, Go makes a copy of the struct and then passes it
// to the function, it doesn't modify the actual struct as the function would only work on the
// copy of that struct valid within the execution lifecycle of the function. After then it's moved
// away from memory. To pass in the main struct we have to pass in a type of the pointer to the person in the args like so
// *person. then wherever the function is called we pass in the address of the person variable using the
// ampersand like so modifyPersonShortName(&person), now modifications done to person would persist outside the function scope
func modifyPersonShortName(person *Person) {
	person.shortName = "chukwu"
	fmt.Println("This is the new shortname within the function scope", person.shortName)
}
