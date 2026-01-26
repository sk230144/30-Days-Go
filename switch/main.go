package main

import "fmt"

func main() {
	day := 4

	switch day {
	case 1, 4:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("other day")
	}

}
