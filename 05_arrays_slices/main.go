package main

import "fmt"

func main() {
	//arrays
	var fruitArr [2]string

	//assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//Declare and Assign at the same time

	vegArr := [2]string{"Carrot", "Celery"}

	fmt.Println(vegArr)
	fmt.Println(vegArr[1])

	vegSlice := []string{"Carrot", "Celery", "Brocollayli", "Brussels Sprouts"}

	fmt.Println(len(vegSlice))
	fmt.Println(vegSlice[2])

	fmt.Println("vegSlice Range:", vegSlice[1:3])
}
