package main

import (
	"fmt"
)

func main() {
	// Arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr) // [1, 2, 3, 4, 5]

	// Slices
	slice := []int{2, 4, 6}
	slice = append(slice, 4) // Add an element
	fmt.Println(slice)       // [2, 4, 6, 4]
}
