package main
import (
	"fmt"
)
func main() {
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it!")
	}
}