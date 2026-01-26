package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	counter := 0;

	for{
		counter++;
		fmt.Println(counter)

		if counter == 3 {
			break
		}
	}
}