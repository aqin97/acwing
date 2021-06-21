package main

import "errors"

var queue []int

//初始化
func Init() {
	queue = make([]int, 0)
}

//判空
func IsEmpty() bool {
	return len(queue) == 0
}

//入队
func EnQueue(x int) {
	queue = append(queue, x)
}

//出队
func DeQueue() (int, error) {
	if IsEmpty() {
		return 0, errors.New("队列为空")
	}

	x := queue[0]
	queue = queue[1:]

	return x, nil
}

//队头
func Head() (int, error) {
	if IsEmpty() {
		return 0, errors.New("队列为空")
	}
	return queue[0], nil
}

func main() {

}
