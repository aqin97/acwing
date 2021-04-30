package main

import (
	"fmt"
	"unicode"
)

var (
	stk_i []int
	stk_c []rune
)

func eval() {
	if len(stk_i) > 1 && len(stk_c) > 0 {
		num1 := stk_i[len(stk_i)-1]
		num2 := stk_i[len(stk_i)-2]
		stk_i = stk_i[:len(stk_i)-2]

		op := stk_c[len(stk_c)-1]
		stk_c = stk_c[:len(stk_c)-1]

		switch op {
		case '+':
			stk_i = append(stk_i, num1+num2)
		case '-':
			stk_i = append(stk_i, num2-num1)
		case '*':
			stk_i = append(stk_i, num1*num2)
		case '/':
			stk_i = append(stk_i, num2/num1)
		}

	}
}

func main() {
	pr := map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2}

	//var input string
	//fmt.Scanf("%s", &input)
	var input string = "2+10"
	inputs := []rune(input)
	for i, data := range inputs {
		//这里的data是rune类型（int32），四字节
		if unicode.IsDigit(data) {
			//err为空，说明当前字符串对应的字符为数字
			x := 0
			j := i
			for j < len(inputs) && unicode.IsDigit(inputs[j]) {
				x = x*10 + int(input[j]-'0')
				j++
			}
			i = j - 1
			fmt.Println(stk_i)
			stk_i = append(stk_i, x)
			fmt.Println(stk_i)
		} else if data == '(' {
			stk_c = append(stk_c, data)
		} else if data == ')' {
			for stk_c[len(stk_c)-1] != '(' {
				eval()
			}
			stk_c = stk_c[:len(stk_c)-1]
		} else {
			for len(stk_c) != 0 && pr[stk_c[len(stk_c)-1]] >= pr[data] {
				eval()
			}
			stk_c = append(stk_c, data)
		}
	}

	for len(stk_c) != 0 {
		eval()
	}

	fmt.Println(stk_i[0])
}
