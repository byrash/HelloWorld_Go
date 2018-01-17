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
	jimPointer := &jim
	// fmt.Println(jimPointer)
	jimPointer.updateName("Jimmy")
	// jim.updateName("test")
	jim.print()

	mySlice := []string{"test", "again"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

/*
Slices can be changed without pointers
*/
func updateSlice(s []string) {
	s[0] = "test changed"
}

//Non pointer version which wont work
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }
func (p *person) updateName(newFirstName string) {
	// (*p).firstName = newFirstName
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
