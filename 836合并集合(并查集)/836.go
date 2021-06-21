package main

import "fmt"

var n, m int
var p []int
var op byte
var a, b int

//并查集的核心操作：寻找祖宗节点，也就是所在集合编号 + 路径压缩
func Find(x int) int {
	if p[x] != x {
		p[x] = Find(p[x])
	}
	return p[x]
}

func main() {
	fmt.Scanf("%d %d\n", &n, &m)
	p = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%c", &op)
		switch op {
		case 'M':
			p[Find(a)] = Find(b)
		case 'Q':
			fmt.Println(Find(a) == Find(b))
		}
	}
}
