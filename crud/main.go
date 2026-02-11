package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
)

type TODO struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool    `json:"completed"`
}

func main() {
	fmt.Println("learing crud")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error get", err)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response:", res.Status)
		return
	}

	defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)   this is normal way of fetch data
	// if err != nil{
	// 	fmt.Println("error", err)
	// }
	// fmt.Println("Real data is", string(data))

	var todo TODO //real world data saving
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error in decoding", err)
		return
	}
	fmt.Println("ToDo", todo)
}
