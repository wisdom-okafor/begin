package main
import (
	"fmt"
)
func main() {
	switch 9 {
		case 1, 5, 10:
			fmt.Print("one, five, ten")
		case 2, 4, 6:
			fmt.Println("two, four, six")
		default:
			fmt.Print("another number")

	
		}
	}