package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}
	fmt.Println("Hello, World!!!!", ids)
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Sum of ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with Map

	phoneNumbers := map[string]int{"Bob": 3035796555, "Sharon": 3036801750, "Mike": 3037863686}

	for k, v := range phoneNumbers {
		fmt.Printf("%s: %v\n", k, v)
	}

	// just get keys

	for k := range phoneNumbers {
		fmt.Println("name " + k)
	}
}
