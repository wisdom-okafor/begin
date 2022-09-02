package main

import "fmt"

func subtraction(x int, y int) int {
	var result int
	result = x - y
	return result

}
func main () {
	var result = subtraction(690000,50000)
	
	fmt.Println(result)
}