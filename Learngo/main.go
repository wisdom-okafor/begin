package main
import (
	"fmt"
	"math"
)
func main() {
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")

	} else {
		fmt.Println("Thes are different")
		}
	}