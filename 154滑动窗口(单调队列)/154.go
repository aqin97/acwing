package main

import (
	"fmt"
)

var n, k int
var a, q []int

func IsEmpty() bool {
	return len(q) == 0
}

func main() {
	fmt.Scanf("%d%d\n", &n, &k)
	a = make([]int, n)
	q = make([]int, 0)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("\n")

	for i := 0; i < n; i++ {
		if !IsEmpty() && i-k+1 > q[0] {
			q = q[1:]
		}
		for !IsEmpty() && a[q[len(q)-1]] >= a[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-k+1 >= 0 {
			fmt.Printf("%d ", a[q[0]])
		}
	}
}
