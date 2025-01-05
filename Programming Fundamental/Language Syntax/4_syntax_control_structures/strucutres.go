package main

import "fmt"

func main() {
	// If
	fmt.Println("If")
	x := 10
	if x > 10 {
		fmt.Println("X is greater than 10")
	} else if x == 10 {
		fmt.Println("X is equal to 10")
	} else {
		fmt.Println("X is less than 10")
	}

	// Switch
	fmt.Println("\nSwitch")
	day := 6

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Another Day")
	}

	// For
	fmt.Println("\nFor")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While
	fmt.Println("\nWhile")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
}
