package main

import (
	"fmt"
)	
func main() {
         var students [5]string
		 fmt.Printf("Students: %v\n", students)
		 students[0] = "Lisa"
		 students[2] = "Ahmed"
		 students[1] = "Arnold"
		 fmt.Printf("Students #1: %v\n", students [1])
		 fmt.Printf("Number of students: %v\n", len(students))	
  
}