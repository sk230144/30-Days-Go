package main

import "fmt"

func addNum(a, b int) int {
	fmt.Println("SSimple function called")
	return a + b
}

func main() {
	fmt.Println("Learning Function")
    add := addNum(3, 3);
	fmt.Println("addition is", add)
}
