package main

import "fmt"

var n, m, q int
var a, b [][]int

func Insert(x1, y1, x2, y2, c int) {
	b[x1][y1] += c
	b[x2 + 1][y1] -= c
	b[x1][y2 + 1] -= c
	b[x2 + 1][y2 + 1] += c
}

func main()  {
	fmt.Scanf("%d%d%d\n", &n, &m, &q)

	a = make([][]int, n + 2)
	for i := 0; i < n + 2; i++{
		a[i] = make([]int, m + 2)
	}
	//fmt.Println(a)
	b = make([][]int, n + 2)
	for i := 0; i < n + 2; i++{
		b[i] = make([]int, m + 2)
	}

	//输入初始数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++{
			fmt.Scanf("%d", a[i][j])
		}
		fmt.Scanf("\n")
	}

	//初始化差分数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++{
			Insert(i, j, i, j, a[i][j])
		}
	}
	/*
	var x1, y1, x2, y2, c int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d%d%d%d%d\n", &x1, &y1, &x2, &y2, &c)
		Insert(x1, y1, x2, y2, c)
	}*/

	//求数组b的前缀和 也就是新的a
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++{
			a[i][j] = b[i][j] + a[i - 1][j] + a[i][j - 1] - a[i - 1][j - 1]
		}
	}

	//最后输出
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++{
			fmt.Printf("%d ", a[i][j])		
		}
		fmt.Printf("\n")
	}
}