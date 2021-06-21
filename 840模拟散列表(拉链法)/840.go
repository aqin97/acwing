package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

var hashTable []*ListNode

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	hashTable = make([]*ListNode, n)

	for i := 0; i < m; i++ {
		var op byte
		var x int
		switch op {
		case 'I':
			Insert(x)
		case 'Q':
			fmt.Println(Query(x))
		}
	}
}

func Hash(x int) int {

	return 0
}

func Insert(x int) {
	if Query(x) {
		return
	}
	h := Hash(x)
	node := new(ListNode)
	node.next = hashTable[h].next
	hashTable[h].next = node
}

func Query(x int) bool {
	h := Hash(x)
	if hashTable[h] == nil {
		return false
	}
	i := hashTable[h].next
	for i != nil {
		if i.val == x {
			return true
		}
		i = i.next
	}
	return false
}
