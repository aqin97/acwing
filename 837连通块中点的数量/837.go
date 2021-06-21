package main

import "fmt"

var n, m int
var p, size []int
var op string
var a, b int

func Find(x int) int {
	if p[x] != x {
		p[x] = Find(p[x])
	}
	return p[x]
}

func main() {
	fmt.Scanf("%d %d\n", &n, &m)
	p = make([]int, n+1)
	size = make([]int, n+1)

	for i := 0; i < n+1; i++ {
		p[i] = i
		size[i] = 1
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%s\n", &op)
		if op == "C" {
			fmt.Scanf("%d %d\n", &a, &b)
			if Find(a) == Find(b) {
				continue
			}
			size[Find(b)] += size[Find(a)]
			p[Find(a)] = Find(b)

		} else if op == "Q1" {
			fmt.Scanf("%d %d\n", &a, &b)
			fmt.Println(Find(a) == Find(b))
		} else if op == "Q2" {
			fmt.Scanf("%d\n", &a)
			fmt.Printf("%d\n", size[Find(a)])
		}
	}
}
