package main

import (
	"fmt"
)	
func main() {
	var students [3]string
	fmt.Printf("students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("students: %v\n", students)
  
}