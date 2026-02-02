package main

import (
	"fmt"
	"strings"
)

func main() {
	//splits based on ,
	data := "apple, orange, bannana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	//count repetetion
	str := "one two three two two"
	count := strings.Count(str, "two")
	fmt.Println(count)

	//space remove
	str = "                hi, me              "
	new := strings.TrimSpace(str)
	fmt.Println(new)

	str1 := "saurabh"
	str2 := "tiwari"
	res := strings.Join([]string{str1, str2}, ",") //.join take array of strings to join , and seconf seprater
	fmt.Println(res, "Joined string")
}