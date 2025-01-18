package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "世界")
	fmt.Println(a, b) // 世界 Hello
}
