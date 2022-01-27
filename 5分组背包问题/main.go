package main

import "fmt"

//N组， 背包体积为V
var N, V int

//每组的个数
var s [100]int

//v[i][j] 表示第i组第j个物品， w同理， f状态数组
var v, w, f [100][100]int

func main() {
	fmt.Scanf("%d %d\n", &N, &V)
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d\n", s[i])
		for j := 0; j <= s[i]; j++ {
			fmt.Scanf("%d %d\n", &v[i][j], &w[i][j])
		}
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= V; j++ {
			for k := 0; k < s[i]; k++ {
				f[i][j] = f[i-1][j]
				if v[i][k] <= j {
					f[i][j] = max(f[i][j], f[i-1][j-v[i][k]]+w[i][k])
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
