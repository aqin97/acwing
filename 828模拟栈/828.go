package main

import "errors"

var stk []int

//初始化
func StackInit() {
	stk = make([]int, 0)
}

//isEmpty
func IsEmpty() bool {
	return len(stk) == 0
}

//push
func Push(x int) {
	stk = append(stk, x)
}

//pop
func Pop() (int, error) {
	if IsEmpty() {
		return 0, errors.New("栈是空的")
	}
	x := stk[len(stk)-1]
	stk = stk[:len(stk)-1]

	return x, nil
}

//top
func Top() (int, error) {
	if IsEmpty() {
		return 0, errors.New("栈是空的")
	}
	return stk[len(stk)-1], nil
}

func main() {

}
