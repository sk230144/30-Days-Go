package main

import "fmt"

func main() {
	var num int
	num = 2
	var ptr *int     // here int say like ptr store int type of data store
	ptr = &num    

	fmt.Println("pointer contains",ptr)

	// getting data using ptr
	fmt.Println("Data get", *ptr)

	//Pointer confirts what operation we applied will do on real value
    *ptr = *ptr + 1000
	fmt.Println("operation on num using pointer", num)

	//default pointer valie is nil
	var pointer *int
	fmt.Println(pointer, "default data")
}