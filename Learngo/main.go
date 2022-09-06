package main
import (
	"fmt" 
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

func main() { 
	aDoctor:= Doctor{
		number: 3,
		actorName: "jon pertwee",
		companions: []string {
			"Liz shaw",
			"Jo Grant",
			
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
}	