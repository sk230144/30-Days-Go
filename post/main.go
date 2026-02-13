package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TODO struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performCreate() {
	todo := TODO{
		UserId:    23,
		Title:     "Saurabh",
		Completed: false,
	}

	//consvert todo i json using marshall
	jsonData, err := json.Marshal(todo)
	if err != nil{
		fmt.Println("error happening", err)
		return
	}

	//convert data in string
	jsonstring := string(jsonData);


	//convert json data in to reader
	jsonReader := strings.NewReader(jsonstring)
    
	myUrl :="https://jsonplaceholder.typicode.com/todos"
	//http post take three things:- url, type of content, reader type data

	//send post request
	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err!=nil{
		fmt.Println("error happnes", err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Response", string(data))
}

func main() {
	fmt.Println("Creating Record")
	performCreate()
}
