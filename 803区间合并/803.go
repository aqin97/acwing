package main

import "fmt"

var n int

type Section struct {
	l int
	r int
}

var a, s []Section

func Swap(a, b *Section)  {
	var tmp Section
	tmp.l = a.l
	tmp.r = a.r
	a.l = b.l
	a.r = b.r
	b.l = tmp.l
	b.r = tmp.r
}

func QuickSort(a []Section, l, r int)  {
	if l >= r {
		return
	}
	x := a[(l + r) / 2].l
	i, j := l - 1, r + 1

	for i < j {
		i++
		for a[i].l < x {
			i++
		}
		j--
		for a[j].l > x {
			j--
		}
		if i < j {
			Swap(&a[i], &a[j])
		}
	}

	QuickSort(a, l, j)
	QuickSort(a, j + 1, r)
}

func main()  {
	fmt.Scanf("%d\n", &n)
	a = make([]Section, n)
	//fmt.Println(a)
	for i := 0; i < n; i++{
		fmt.Scanf("%d %d\n", &a[i].l, &a[i].r)
		//fmt.Println(a)
	}
	//对初试数据进行排序（按区间左端点大小排序
	QuickSort(a, 0, n - 1)
	//fmt.Println(a)

	st, ed := a[0].l, a[0].r

	//遍历排好序的a
	for i := 1; i < n; i++ {
		if a[i].l <= ed && a[i].r > ed {
			ed = a[i].r
		} else if ed < a[i].l {
			res := Section{st, ed}
			s = append(s, res)
			st = a[i].l
			ed = a[i].r
		}
	}
	res := Section{st, ed}
	s = append(s, res)
	//fmt.Println(s)
	fmt.Println(len(s))
}