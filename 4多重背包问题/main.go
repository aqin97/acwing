package main

import "fmt"

var N, V int

var v, w, s []int

var f [][]int

func main() {
	fmt.Scanf("%d %d\n", &N, &V)
	v = make([]int, N+1)
	w = make([]int, N+1)
	s = make([]int, N+1)
	f = make([][]int, N+1)
	for i := 0; i <= N; i++ {
		f[i] = make([]int, V+1)
	}
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d %d %d\n", &v[i], &w[i], &s[i])
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= V; j++ {
			for k := 0; k*v[i] <= j && k <= s[i]; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}
	fmt.Println(f)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
