package main

import "fmt"

func main() {
	fmt.Println("Hello, World!!!!")
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("a: %T b: %T\n", a, b)
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change Value with pointer

	*b = 10
	fmt.Println(a)

}
