package main 

import (
    "fmt"
)

var pl = fmt.Println

func main() {
// %d : Integer
// %s : String
// %t : Boolean
// %f : Float
// %c : Character
// %o : Base 8
// %x : Base 16
// %T : Type of supplied value
// %v : Guesses based on data type

fmt.Printf("%s, %d, %t, %f, %c, %o, %x\n",
"stuff", 12, true, 1.12, 'k', 1, 1)
fmt.Printf("%9f\n", 3.14)
fmt.Printf("%.2f\n", 3.141592)
fmt.Printf("%.2f\n", 3.141592)

sp1 := fmt.Sprintf("%9.f\n", 3.141592)
pl(sp1)
}