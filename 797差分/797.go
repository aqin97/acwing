package main

import "fmt"

var n, m int
var a, b []int

func Insert(l, r, c int) {
	if len(a) == r + 1 {
		b[l] += c
	} else {
		b[l] += c
		b[r + 1] -= c
	}
}

func main()  {
	fmt.Scanf("%d%d\n", &n, &m)
	a = make([]int, n + 1)
	b = make([]int, n + 1)

	//输入初始数组
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Scanf("\n")
	//构建差分数组
	for i := 1; i <= n; i++ {
		Insert(i, i, a[i])
	}
	fmt.Println(b)
	var l, r, c int
	//对差分数组进行变换
	for i := 0; i < m; i++ {
		fmt.Scanf("%d%d%d\n", &l, &r, &c)
		Insert(l, r, c)
		fmt.Println(b)
	}

	for i := 1; i <= n; i++ {
		b[i] += b[i - 1]
	}

	//输出变化过后的原始数组
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", b[i])
	}
}