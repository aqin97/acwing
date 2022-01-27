package main

import "fmt"

var INF int = 1e9

//三角形的层数
var N int

//三角形数组和状态数组
var tri, f [][]int

func main() {
	fmt.Scanf("%d\n", &N)
	tri = make([][]int, N+1)
	f = make([][]int, N+1)
	for i := 0; i <= N; i++ {
		tri[i] = make([]int, N+1)
		f[i] = make([]int, N+1)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scanf("%d", &tri[i][j])
		}
		fmt.Scanf("\n")
	}
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			f[i][j] = -INF
		}
	}
	for i := 0; i <= N; i++ {
		fmt.Println(tri[i])
	}
	res := -INF
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			f[i][j] = max(f[i-1][j-1]+tri[i][j], f[i-1][j]+tri[i][j])
			res = max(res, f[i][j])
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
