package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("file.tx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
