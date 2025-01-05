package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Go routine!")
}

func main() {
	go sayHello()

	time.Sleep(5 * time.Millisecond)
	fmt.Println("Bye sayHello() <_<!")
}
