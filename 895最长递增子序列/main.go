package main

import (
	"fmt"
)

var n int
var a, dp []int

func main() {
	fmt.Scanf("%d\n", &n)
	a = make([]int, n)
	dp = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
