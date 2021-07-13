package main

import (
	"fmt"
	"sort"
)

type citizen struct {
	id int
	order int
	candy int64
}

func main() {
	var n int
	var k int64
	fmt.Scanf("%d %d", &n, &k)
	citizens := make([]citizen, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)
		citizens[i] = citizen {
			id: x,
			order: i,
			candy: 0,
		}
	}
	sort.Slice(citizens, func(i, j int) bool {
		return citizens[i].id < citizens[j].id
	})
	at_least_candies := k / int64(n)
	remain_candies := k % int64(n)
	for i := 0; i < n; i++ {
		citizens[i].candy = at_least_candies
		if int64(i) < remain_candies {
			citizens[i].candy++
		}
	}
	sort.Slice(citizens, func(i, j int) bool {
		return citizens[i].order < citizens[j].order
	})
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", citizens[i].candy)
	}
}