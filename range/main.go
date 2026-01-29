package main

import "fmt"


// range willl give us value and index
func main() {
	numbers := []int{7, 22, 33, 14}

	for index, value := range numbers {
		fmt.Printf("Index of numbers is %d, Value of numbers is %d\n", index, value)
	}

	data := "hellow world"

	for index, value := range data{
		fmt.Printf("Index of data is %d, Value of data is %c\n", index, value)
	}
}