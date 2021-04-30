package main

import "fmt"

var n int
var e, l, r []int
var head, tail, idx int

func InitList() {
	head = 0
	tail = 1
	r[head] = tail
	l[tail] = head
	idx = 2
}

func InsertL(k, x int) {
	e[idx] = x
	r[idx] = r[k]
	l[idx] = k
	l[r[k]] = idx
	r[k] = idx
	idx++
}

func InsertR(k, x int) {
	InsertL(l[k], x)
}

func Delete(k int) {
	l[r[k]] = l[k]
	r[l[k]] = r[k]
}

func main() {
	fmt.Scanf("%d\n", &n)

}
