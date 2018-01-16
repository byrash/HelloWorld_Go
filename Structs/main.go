package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// alex := person{ "Alex",  "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "test@email.com"
	// alex.contact.zipCode = 1234
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "him@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
