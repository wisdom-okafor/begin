package main
import (
	"fmt" 
)
func main() { 
	statepopulations := make(map[string]int)
	statepopulations = map[string]int{
		       "california":     39250017,
			   "Texas":          27862596,
			   "Florida":        20612439,
			   "New york":       19745289,
			   "Pennsylvania":   12802503,
			   "Illinois":       12801539,
			   "Ohio":           11614373,
}
fmt.Println(statepopulations)
statepopulations["Georgia"] = 10310371
fmt.Println(statepopulations)
}