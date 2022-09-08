package main
import (
	"fmt"
)
func main() {
	s := "hello Go"
	for k, v := range s {
		fmt.Println(k, v)
	}
}