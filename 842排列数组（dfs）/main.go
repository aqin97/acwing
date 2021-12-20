package main

import "fmt"

//input
var n int

//result
var path []int

//dfs status
var seen []bool

func main() {
	fmt.Scanf("%d\n", &n)
	path = make([]int, n+1)
	seen = make([]bool, n+1)

	dfs(0)
}

func dfs(u int) {
	//u means current search layer
	//border :dfs reach the last layer,thus u == n
	if u == n {
		for i := 1; i < n; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Printf("%d\n", path[n])
		return
	}

	for i := 1; i <= n; i++ {
		//status check
		if !seen[i] {
			path[u+1] = i
			seen[i] = true
			dfs(u + 1)
			//status recovery
			seen[i] = false
		}
	}
}
