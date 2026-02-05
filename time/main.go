package main

import (
	"fmt"
	"time"
)

func main(){
	 currentTime := time.Now()
	 fmt.Println("time", currentTime)

	 formattedData := currentTime.Format("02-01-2006, Monday, 15:04:05") //dd-MM-YYYY not written here always define data when go launched
	 ampmFormat := currentTime.Format("3:04 PM") //always write this for am pm
	 fmt.Println("formatted time", formattedData, ampmFormat)


	 //add one day in cuurent time
	 new_date := currentTime.Add(24 * time.Hour).Format("2006/01/02 Monday")
	 fmt.Println(new_date, "formatted new date")
}