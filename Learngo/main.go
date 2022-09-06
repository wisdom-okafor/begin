package main
import (
	"fmt" 
)

func main() { 
	aDoctor:= struct{name string}{name: "John Pertwee"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)	
}	