package shape

import "fmt"

type Rectangle struct {
	Height uint64
	Width  uint64
}

func (r Rectangle) Draw() {
	fmt.Println("Rectangle is drawing")
}
