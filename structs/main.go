package main

import "fmt"

// Contact struct
type contactInfo struct {
	email   string
	zipCode int
}

// Creating the person structure to serve as the blueprint of the type person.
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	sam := person{
		firstName: "Samuel",
		lastName:  "Ofori",
		contactInfo: contactInfo{
			email:   "oforisml@gmail",
			zipCode: 23456,
		},
	}
	
	//	&variable gets the address of the variable
	sam.updateName("Jim")
	sam.print()

}
	// *variable 	gives as the value at variable.
	// When * is in front of a type, it means we are working with a pointer
func (pointerToPerson *person) updateName(newFirstName string) {

	//When the * is infront of a pointer, it means we want to manipulate the pointer's value.
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
