package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
)

type TODO struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func readFunction() {
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

func createFunction() {
	todo := TODO{
		UserId:    23,
		Title:     "Saurabh",
		Completed: false,
	}

	//consvert todo i json using marshall
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error happening", err)
		return
	}

	//convert data in string
	jsonstring := string(jsonData)

	//convert json data in to reader
	jsonReader := strings.NewReader(jsonstring)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	//http post take three things:- url, type of content, reader type data

	//send post request
	res, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("error happnes", err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Response", string(data))
}

func updateFunction() {
	todo := TODO{
		UserId:    23,
		Title:     "Saurabh Tiwari",
		Completed: false,
	}

	//consvert todo i json using marshall
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error happening", err)
		return
	}
	//convert data in string
	jsonstring := string(jsonData)

	//convert json data in to reader
	jsonReader := strings.NewReader(jsonstring)

	myUrl := "https://jsonplaceholder.typicode.com/todos/2"

	//Create put method

	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)

	if err != nil {
		fmt.Println("error in req", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	//Send the request
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("error here in request", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Response during update", string(data), res.StatusCode)
}

func main() {
	// readFunction()
	// createFunction()
	updateFunction()
}
