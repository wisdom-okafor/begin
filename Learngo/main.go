package main

import (
	"fmt"
)	
func main() {
        a := []int{1, 2, 3}
		b := a 
		b[1] = 5
		fmt.Println(a)
		fmt.Println(b)
		fmt.Printf("Lenght: %v\n", len(a))
		fmt.Printf("capacity: %v\n", cap(a)) 
}