package main

import (
	"fmt"
	"reflect"
	"strconv"

	
)

var pl = fmt.Println

func main() {
	cV3 := "50000000"
    cV4, err := strconv.Atoi(cV3)
	  pl(cV4, err, reflect.TypeOf(cV4))
	
}