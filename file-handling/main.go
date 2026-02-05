package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error to create")
		return
	}
	defer file.Close() // after creating file make resources free

	content := "Hii saurabh"
	_, error := io.WriteString(file, content)

	if(error != nil){
		fmt.Println("Error in write file:", error)
		return
	}

	fmt.Println("File created succesfully")
}