package main 

import (
    "fmt"
)

var pl = fmt.Println

type Animal interface{
    AngrySound()
    HappySound()

}
type Cat string

func (c Cat) Attack(){
    pl("cat Attacks it's prey")

}
func (c Cat) Name() string{
    return string(c)
}
func (c Cat) AngrySound(){
}
func (c Cat) HappySound() {
    pl("cat says purrrrrr")
}

func main() {
    var kitty Animal
    kitty = Cat("kitty")
    kitty.AngrySound()
    var kitty2 Cat = kitty.(Cat)
    kitty2.Attack()
    pl("Cats Name :", kitty2.Name())

}