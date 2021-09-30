package main

import (
	"fmt"
	"hieu_test/shape"
)

func main() {
	fmt.Print("hello em")
	myShape := shape.Shape{shape.Rectangle{10, 5}, 17}
	fmt.Println("myShape = ", myShape)
	// fmt.Println("height = ", myShape.Height)
	// fmt.Println("height = ", myShape.Rectangle.Height)
	shape.Execute(myShape)
	myRec := shape.Rectangle{20, 10}
	shape.Execute(myRec)
}
