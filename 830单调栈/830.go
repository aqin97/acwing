package main

import (
	"fmt"
)

var n int
var a []int
var stk []int

func IsEmpty() bool {
	return len(stk) == 0
}

func main() {
	fmt.Scanf("%d\n", &n)
	a = make([]int, n)
	stk = make([]int, 0)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("\n")

	for i := 0; i < n; i++ {
		for !IsEmpty() && stk[len(stk)-1] >= a[i] {
			stk = stk[:len(stk)-1]
		}
		if IsEmpty() {
			fmt.Printf("-1 ")
		} else {
			fmt.Printf("%d ", stk[len(stk)-1])
		}
		stk = append(stk, a[i])
	}
}
