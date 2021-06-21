package main

import "fmt"

func GetNext(p string) []int {
	next := make([]int, len(p)+1)
	next[0] = -1
	j, k := 0, -1
	for j < len(p)-1 {
		if k == -1 || p[j] == p[k] {
			j++
			k++
			if p[j] == p[k] {
				next[j] = next[k]
			} else {
				next[j] = k
			}
		} else {
			k = next[k]
		}
	}
	return next
}

func main() {
	var s, p string
	fmt.Scanf("%s\n%s\n", &s, &p)
	next := GetNext(p)
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(p) {
		fmt.Printf("%d %d\n", i-j, i-j-1+len(p))
	} else {
		fmt.Println(-1)
	}
}
