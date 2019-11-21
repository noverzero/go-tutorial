package main

import (
	"fmt"
	"strconv"
)

//define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting method value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " I am " + strconv.Itoa(p.age)
}

// hasBirthday method pointer receiver

func (p *Person) hasBirthday() {
	p.age++
}

//getMarried pointer value receiver
func (p *Person) getMarried(spouse Person) {
	spouseLastName := spouse.lastName
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//init person using struct

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 24}

	//Alternative syntax where keys are inferred based on sequence

	person2 := Person{"Doug", "Jarvis", "Lakewood", "m", 43}

	fmt.Println(person1)
	fmt.Println(person2)

	//Get single field

	person1.age++

	fmt.Println(person1.age)
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried(person2)
	person2.getMarried(person1)
	fmt.Println(person1.greet())

}
