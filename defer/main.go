package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("start of program")
	defer fmt.Println("mid of program") //defer prints data after all program done
	fmt.Println("end of program")

	data := add(5, 7)
    defer fmt.Println("Sum of data", data); //when 2 defer present then defers push in stack, print like last in first out
}