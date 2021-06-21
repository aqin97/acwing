package main

import "fmt"

var n, m, idx int
var cnt []int
var son [][]int
var op byte
var str string

func Insert(x string) {
	chs := []byte(x)
	var p int
	for _, ch := range chs {
		u := int(ch - 'a')
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	cnt[p]++
}

func Qusery(x string) int {
	chs := []byte(x)
	var p int
	for _, ch := range chs {
		u := int(ch - 'a')
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	return cnt[p]
}

func main() {
	fmt.Scanf("%d %d\n", &n, &m)
	cnt = make([]int, n)
	son = make([][]int, n)
	for i := 0; i < n; i++ {
		son[i] = make([]int, 26)
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%c %s\n", &op, &str)
		switch op {
		case 'I':
			Insert(str)
		case 'Q':
			fmt.Printf("%d", Qusery(str))
		}
	}
}
