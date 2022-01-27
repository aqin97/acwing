package main

import (
	"container/list"
	"fmt"
)

func main() {
	var stk list.List
	stk.PushFront(1)
	stk.PushFront(2)
	for e := stk.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
