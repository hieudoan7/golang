package main
import "fmt"

func main(){
	var p int
	fmt.Scanf("%d", &p)
	n_coins := 0
	factorial := 1
	for i:=2; i<12; i++ {
		remainder := p % (factorial*i)
		n_coins += remainder / factorial
		factorial *= i
		p -= remainder
	}
	fmt.Println(n_coins)
}