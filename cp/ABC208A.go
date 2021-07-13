package main
import "fmt"

func main(){
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a <= b && b <= 6*a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}