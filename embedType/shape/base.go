package shape

type Object interface {
	Draw()
}

type Shape struct {
	Rec Rectangle
	Num uint64
}

// func (s Shape) Draw() {
// 	fmt.Println("Shape is drawing")
// }

func Execute(obj Object) {
	obj.Draw()
}
