package main

import "fmt"

//n items, m volume
var n, m int

//vol and wealth of item
var v, w []int

//result
var f [][]int

func main() {
	fmt.Scanf("%d %d\n", &n, &m)
	v = make([]int, n+1)
	w = make([]int, n+1)

	f = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
	}
	//fmt.Println(f)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d\n", &v[i], &w[i])
	}
	//f[0][0~m] == 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i-1][j-v[i]]+w[i])
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
