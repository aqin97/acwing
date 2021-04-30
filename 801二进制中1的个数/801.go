package main

import "fmt"

var n int
var a, s []int

func lowbit(n int) int {
	return n & -n
	//return n & (^n + 1)
}

func Calc(n int) int {
	var cnt int

	for n != 0 {
		n -= lowbit(n)
		cnt++
	}
	return cnt
}

func main()  {
	fmt.Scanf("%d\n", &n)
	a = make([]int, n)
	s = make([]int, n)
	
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("\n")
	
	for i := 0; i < n; i++ {
		s[i] = Calc(a[i])
		fmt.Printf("%d ", s[i])
	}
}