package main

import "fmt"

var n int
var a, s []int

func max(a, b int) int {
	if a > b {
		return a
	} 
	return b
}

func main()  {
	fmt.Scanf("%d\n", &n)
	a = make([]int, n)
	s = make([]int, 100010)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("\n")

	var res int = 0
	for i, j := 0, 0; i < n; i++ {
		s[a[i]]++
		for s[a[i]] > 1 {
			s[a[j]]--
			j++
		}
		res = max(res, i - j + 1)
	}

	fmt.Printf("%d\n", res)
}