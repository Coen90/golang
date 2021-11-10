package main

import "fmt"

func sliceAndArray() {
	// array and how to print array
	// foods := [3]string{"potato", "pizza", "tasta"}
	// for _, food := range foods {
	// 	fmt.Println(food)
	// }
	// for i := 0; i < len(foods); i++ {
	// 	fmt.Println(foods[i])
	// }

	// slice
	// foods := []string{"potato", "pizza", "tasta"}
	// for i := 0; i < len(foods); i++ {
	// 	fmt.Println(foods[i])
	// }

	foods := []string{"potato", "pizza", "tasta"}
	fmt.Printf("%v\n", foods)
	fmt.Println(len(foods))
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)
	fmt.Println(len(foods))

}