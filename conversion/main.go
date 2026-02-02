package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 32
	fmt.Printf("Type of num is %T\n", num)

    //Number to float conversion
	var data float64 = float64(num)
	data += 1.23
	fmt.Println("data is", data)
	fmt.Printf("Type of data is %T\n", data)

    //Number to string conversion
	str := strconv.Itoa(num)
	fmt.Println("data is", str)
	fmt.Printf("Type of str is %T\n", str)
   
	//String to number
	number_int := "1234"
	number_num, _ := strconv.Atoi(number_int) // here strconv.Atoi return number and error, so instead of get error we denote _
	fmt.Println("data is", number_num)
	fmt.Printf("Type of str is %T\n", number_num)

	//string to float
	number_str := "1234.2"
	number_float, _ := strconv.ParseFloat(number_str, 64) // here strconv.ParseFloat take string and number of bits get
	fmt.Println("data is", number_float)
	fmt.Printf("Type of str is %T\n", number_float)
}	
