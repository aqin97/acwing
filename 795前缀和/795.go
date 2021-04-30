package main

import "fmt"

var n, m, l, r int
var a, s []int

func main() {
	fmt.Scanf("%d%d\n", &n, &m)
	a = make([]int, n+1)
	s = make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("\n")
	fmt.Println(a)
	s[0] = 0
	for i := 1; i < n+1; i++ {
		s[i] = s[i-1] + a[i]
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d\n", &l, &r)
		fmt.Printf("%d\n", s[r]-s[l-1])
	}
}
