package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if(b == 0){
      return  0, fmt.Errorf("denominator must be greater than 0")
	}
	return a/b, nil
}

func main() {
	fmt.Println("starting error handling")
    ans1, _ := divide(10, 0);   // if we dont want to pront error
	fmt.Println(ans1, "ans1")

	ans , err := divide(10, 0)
	if err != nil {
		fmt.Println("error happens")
	}
	fmt.Println(ans, "ans is this")
}