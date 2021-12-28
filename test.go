package main
import (
	"fmt"
)

type student struct {
	name string
	age uint32
}

func main() {
	var a *student
	fmt.Println(a) //nil
	b := &student{}
	fmt.Println(b)  //&{ 0}
}