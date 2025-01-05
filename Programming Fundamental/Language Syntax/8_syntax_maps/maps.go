package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 45
	fmt.Println(m["Alice"], m["Bob"]) // 25 45
}
