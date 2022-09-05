package main

import "fmt"

func mydetails(name string, age int, phonenumber string, height float64) string {

    var result = fmt.Sprintf("My name is %s, i am %d years old, my phonenumber is %s, and my height is %f", name, age, phonenumber, height)
	return result 
}
	func main () {
    var result = mydetails("wisdom okafor", 19, "+2347035392310", 5.7)
	fmt.Println(result)

}

