package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing out loops")

	for i := 0; i < 10; i++ {
		fmt.Println("Occurrence: ", i+1)
		printVal(i)
	}
}

func printVal(a int) {
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unsupported operation")
	}
}
