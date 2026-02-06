package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("learn url")
	myUrl := "https://jsonplaceholder.typicode.com/todos?key=value"
    parsedUrl, err := url.Parse(myUrl)
	if(err != nil){
		fmt.Println("error", err)
	}
	fmt.Printf("Type of myUrl is %T\n", parsedUrl);
    fmt.Println("Scheme", parsedUrl.Scheme)
	fmt.Println("Host", parsedUrl.Host)
	fmt.Println("Path", parsedUrl.Path)
	fmt.Println("Rawquery", parsedUrl.RawQuery)

	//Modify url

	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=saurabh"
	newUrl := parsedUrl.String()
	fmt.Println("new url", newUrl)
}