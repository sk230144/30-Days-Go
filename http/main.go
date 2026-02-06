package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if(err != nil){
		fmt.Println("Error in fetch", err)
		return
	}
	fmt.Printf("Type of data %T\n", res)
	data, err := ioutil.ReadAll(res.Body)  //read response body and give byte slice
	if(err != nil){
		fmt.Println("error in response", err)
	}
	fmt.Println("response:", string(data)) //initially data is in binary,convert data in to string

	defer res.Body.Close() //always free the connection from http after read in go
}