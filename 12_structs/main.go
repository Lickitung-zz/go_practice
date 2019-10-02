package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Brandon", lastName: "Evans", city: "Austin", gender: "M", age: 25}
	person2 := Person{firstName: "Kayla", lastName: "Rausch", city: "Austin", gender: "F", age: 26}
	// person1 := Person{"Brandon", "Evans", "Austin", "M", 25}

	fmt.Println(person1)
	fmt.Println(person1.firstName)

	person1.hasBirthday()
	person2.getMarried("Evans")
	person1.getMarried("Evans")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
