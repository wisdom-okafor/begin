package main

import "fmt"

func divition(x int, y int) int {
	var result int
	result = x / y
	return result

}
func main () {
	var result = divition(1678,12)
	fmt.Println(result)
}