package main

import "fmt"

var m int
var head, idx int
var e, ne []int

func InitList() {
	head = -1
	idx = 0
}

func HeadInsert(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

func Delete(k int) {
	if k == -1 {
		head = ne[head]
	} else {
		ne[k] = ne[ne[k]]
	}
}

func Insert(k, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

func PrintList() {
	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
	fmt.Println()
}

func main() {
	fmt.Scanf("%d\n", &m)

	e = make([]int, m)
	ne = make([]int, m)

	InitList()
	for i := 0; i < m; i++ {
		var op byte
		var x, k int
		fmt.Scanf("%c", &op)

		if op == 'H' {
			fmt.Scanf("%d\n", &x)
			HeadInsert(x)
			//fmt.Println(x, e, ne)
		} else if op == 'D' {
			fmt.Scanf("%d\n", &k)
			Delete(k - 1)
			//fmt.Println(k, e, ne)
		} else if op == 'I' {
			fmt.Scanf("%d %d\n", &k, &x)
			Insert(k-1, x)
			//fmt.Println(k, x, e, ne)
		}
		//PrintList()
	}

	PrintList()
}
