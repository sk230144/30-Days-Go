package main

import "fmt"

func main() {
	fmt.Println("Array learn in go")
    // Hello World 
	var name[5] string;
	name[0] = "saurabh"
	name[1] = "vir"

	var numbers = [8]int{1, 2, 3, 4, 6, 7, 8, 0}

	fmt.Println(name, "length " , len(name), numbers[7])

	// Go always set default value of type in array, like 0, "", false

	var price[10] int;
	fmt.Println(price, "price is here")
}
   