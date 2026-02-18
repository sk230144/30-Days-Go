package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000*time.Millisecond)
	fmt.Println("Function sayhello execution finished")
}

func sayHi() {
	
	time.Sleep(2000*time.Millisecond)
	fmt.Println("Hii")
}

func main() {
	fmt.Println("Learning go routine")
    go sayHello() //that will cuncrrently or paralle;y run, without depend on others
	go sayHi() //we use go keyword for goroutine, for concurrent run

	time.Sleep(3000*time.Millisecond) //it wait every function to execute in 1 secs
}