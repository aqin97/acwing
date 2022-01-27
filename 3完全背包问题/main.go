package main

import "fmt"

//N种物品(编号1—N), V容量的背包
var N, V int

//每种物品的体积v和价值w
var v, w []int

//状态，表示最大的w
var f []int

func main() {
	fmt.Scanf("%d %d\n", &N, &V)
	v = make([]int, N+1)
	w = make([]int, N+1)
	f = make([]int, V+1)
	for i := 1; i < N+1; i++ {
		fmt.Scanf("%d %d\n", &v[i], &w[i])
	}
	/*暴力解
	for i := 1; i < N+1; i++ {
		for j := 0; j <= V; j++ {
			for k := 0; k*v[i] <= j; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}
	*/
	//时间优化解
	/*
		for i := 1; i <= N; i++ {
			for j := 0; j <= V; j++ {
				f[i][j] = f[i-1][j]
				if j >= v[i] {
					f[i][j] = max(f[i][j], f[i][j-v[i]]+w[i])
				}
			}
		}*/
	//空间优化

	for i := 1; i <= N; i++ {
		for j := v[i]; j <= V; j++ {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	fmt.Println(f)
	fmt.Println(f[V])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
