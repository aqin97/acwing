package main

import "fmt"

var n, m, q int
var a, s [][]int

func calc(x1, y1, x2, y2 int) int {
	return s[x2][y2] - s[x1 - 1][y2] - s[x2][y1 - 1] + s[x1 - 1][y1 - 1]
}

func main()  {
	fmt.Scanf("%d%d%d\n", &n, &m, &q)
	a = make([][]int, n + 1)
	for i := 0; i < m; i++ {
		a[i] = make([]int, m + 1)
	}
	//fmt.Println(a)
	s = make([][]int, n + 1)
	for i := 0; i < m; i++ {
		s[i] = make([]int, m + 1)
	}

	//输入原始数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++{
			fmt.Scanf("%d", &a[i][j])
		}
		fmt.Scanf("\n")
	}
	//fmt.Scanf("\n")

	//计算前缀和数组
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] = s[i - 1][j] + s[i][j - 1] - s[i - 1][j - 1] + a[i][j]
		}
	}
	fmt.Println(s)
	var x1, y1, x2, y2 int

	for i := 0; i < q; i++ {
		fmt.Scanf("%d%d%d%d\n", &x1, &y1, &x2, &y2)
		fmt.Printf("%d\n", calc(x1, y1, x2, y2))
	}
}