package main

import "fmt"

const N = 20

var n int
var board [N][N]byte
var col, dg, udg [N]bool

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%c", board[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		if !col[i] && !dg[u+i] && !udg[n-u+i] {
			board[u][i] = 'Q'
			col[i] = true
			dg[u+i] = true
			udg[n-u+i] = true

			dfs(u + 1)
			col[i] = false
			dg[u+i] = false
			udg[n-u+i] = false
			board[u][i] = '.'
		}
	}
}

func main() {
	fmt.Printf("please input an integer less equal than 8:")
	fmt.Scanf("%d\n", &n)
	//init board
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	dfs(0)
}
