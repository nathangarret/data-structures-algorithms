package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func div(x int, y int) int {
	return x / y
}

func mult(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(42, 59))  // 101
	fmt.Println(sub(99, 11))  // 88
	fmt.Println(div(110, 10)) // 11
	fmt.Println(mult(12, 12)) // 144
}
