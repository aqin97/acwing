package main

import "fmt"

//输入
var n int

//结果数组
var path []int

//记录状态
var seen []bool

func dfs(u int) {
	//参数u表示搜索到第几层
	//边界条件，搜索到了最后一层即第n层
	if u == n {
		for i := 0; i < n-1; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println(path[n-1])
		return
	}

	for i := 1; i <= n; i++ {
		if !seen[i] {
			//处理好这一层往下走
			path[u] = i
			seen[i] = true
			dfs(u + 1)
			//恢复现场
			path[u] = 0
			seen[i] = false
		}
	}

}

func main() {
	fmt.Scanf("%d\n", &n)
	seen = make([]bool, n+1)
	path = make([]int, n+1)
	dfs(0)
}
