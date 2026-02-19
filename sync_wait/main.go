package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i, "started")
	fmt.Println(i, "end")
}

func main() {
	// fmt.Println("Learning go routine")

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		//increment worker group counter
		wg.Add(1)
		go worker(i, &wg)
	}

	//wait for all workers to finish then after this codes will work
	wg.Wait()
	fmt.Println("worker task completed")
}
