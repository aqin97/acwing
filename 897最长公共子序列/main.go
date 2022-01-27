package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scanf("%s\n%s\n", &s1, &s2)
	fmt.Println(s1, s2)
	n, m := len(s1), len(s2)
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1])
			if s1[i-1] == s2[j-1] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	for i := 0; i <= n; i++ {
		fmt.Println(f[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
