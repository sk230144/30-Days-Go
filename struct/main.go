package main
import "fmt"

//struct mostly use for making data types

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var prince Person
	prince.FirstName="Ramu"

	fmt.Println("Data", prince)

	person1 := Person{
		FirstName: "Rahu",
		LastName: "Sign",
		Age: 24,
	}
	fmt.Println("Data od man", person1)
}
