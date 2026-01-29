// maps are used to associate values with keys

package main 

import "fmt"

func main(){

	// we want to make string with number key
	// saurabh have 34 mark
	studentGrade := make(map[string]int)

	studentGrade["Saurabh"] = 34
	studentGrade["Virat"] = 37
	studentGrade["Hira"] = 0

	fmt.Println("Fetch Marks of Saurabh: ", studentGrade["Saurabh"]) //fetch
	studentGrade["Hira"] = 10; //update
	delete(studentGrade, "Virat") // delete

   fmt.Println("Marks of Deleted Virat: ", studentGrade["Virat"]) //fetch
   fmt.Println("Marks of Updated Hira: ", studentGrade["Hira"]) //fetch

   // Checking key is present
   // first value, grade return value, second one return presnt or not in true or false
   grades, exists := studentGrade["Raju"]
   fmt.Println("exist or not", exists)
   fmt.Println("grade of this", grades)


   //declare + assign

   persons := map[string]int{
      "Ramu": 56,
	  "bobby":45,
   }

   for index, value := range persons{
	fmt.Printf("Key is %s and value is %d\n", index, value)
   }
}