package main

import "fmt"

func main() {
	var x int = 42
	var p *int = &x
	fmt.Println(*p)
}
