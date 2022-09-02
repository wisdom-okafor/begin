package main

import "fmt"

func addition(x int, y int) int {
	var result int
	result = x + y
	return result
}

func main() {
	var result = addition(9,9)
	fmt.Println(result)
}