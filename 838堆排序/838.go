package main

import "fmt"

var n, m int
var size int
var heap []int

//这里的x表示第x个节点
func Down(x int) {
	t := x
	if 2*x <= size && heap[2*x] < heap[t] {
		t = 2 * x
	}
	if 2*x+1 <= size && heap[2*x+1] < heap[t] {
		t = 2*x + 1
	}
	if t != x {
		heap[t], heap[x] = heap[x], heap[t]
		Down(t)
	}
}

func main() {
	fmt.Scanf("%d %d\n", &n, &m)
	size = n
	heap = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Scanf("%d", &heap[i])
	}

	//建堆
	for i := n / 2; i != 0; i-- {
		Down(i)
	}
	fmt.Println(heap)
	//输出前m个小的数
	for i := 0; i < m; i++ {
		fmt.Printf("%d ", heap[1])
		heap[1] = heap[size]
		size--
		Down(1)
	}
}
