package main
import (
	"fmt"
)
func main() {
	statepopulations := map[string]int{
		"California":     39250017,
		"Texas":          27862596,
		"Florida":        20612439,
		"New york":       19745289,
		"Pennsylvania":   12802503,
		"Illinois":       12801539,
		"Ohio":           11614373,
	}
	for k, v := range statepopulations {
		fmt.Println(k, v)
	}
}