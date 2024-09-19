package main

import (
	"fmt"

	"Adlet/main_function/container"
	"Adlet/main_function/implementation"
)

// import "Adlet/implementation"

func main() {
	var input string = "\\\\"
	// fmt.Println("Initial:", []byte(input))

	implementation.Main_Operation(input, "standard.txt", true)
	fmt.Print(container.Result)
}
