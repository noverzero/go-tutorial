package main

import "fmt"

func main() {
	fmt.Println("Hello, World!!!!")
	//Define Map
	emails := make(map[string]string)

	//assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(len(emails))
	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails["Bob"])
	fmt.Println(emails)

	//Declare Map and add key values
	phoneNumbers := map[string]int{"Bob": 3035796555, "Sharon": 3036801750, "Mike": 3037863686}

	fmt.Println(len(phoneNumbers))
	fmt.Println(phoneNumbers)
	fmt.Println(phoneNumbers["Bob"])

}
