package main

import "fmt"

//输入
var n int

type Node struct {
	Val  int
	Next *Node
}

type Head *Node

var G []Head

func Insert(dst, val int) {
	node := &Node{
		Val: val,
	}
	node.Next = G[dst].Next
	G[dst].Next = node
}

func main() {
	fmt.Scanf("%d\n", &n)
	G = make([]Head, n)
}
