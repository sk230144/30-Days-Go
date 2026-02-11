package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{Name: "Saurab", Age: 34, IsAdult: true}
	fmt.Println(person, "person")

	//convert person in to json
	jsonData, err := json.Marshal(person)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("data in json string", string(jsonData))

	//unmarshaling or decoding the data
    var persondata Person;
	err = json.Unmarshal(jsonData, &persondata)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("data after unmarshal", persondata)

}
