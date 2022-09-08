package main
import (
	"fmt"
)
func main() {
	var i interface{} = [3]int{}
	switch i. (type) {
	case int:
			fmt.Println("i is an int")
	case float64:
			fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3] int:
		fmt.Println("i is [2]int")
	default:
			fmt.Print("i is another type")
		}
	}