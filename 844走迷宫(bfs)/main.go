package main

import "fmt"

const N = 100

//input
var n, m int

//input map
var board [N][N]int

//distance from (0,0)
var distance [N][N]int

type point struct {
	x, y int
}

var queue []point

func main() {
	fmt.Scanf("%d %d\n", &n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &board[i][j])
		}
		fmt.Scanf("\n")
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			distance[i][j] = -1
		}
	}
	distance[0][0] = 0

	fmt.Println(bfs())
}

func bfs() int {
	queue = append(queue, point{0, 0})
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}

	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x, y := p.x+dx[i], p.y+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && board[x][y] == 0 && distance[x][y] == -1 {
				distance[x][y] = distance[p.x][p.y] + 1
				queue = append(queue, point{x, y})
			}
		}
	}

	return distance[n-1][m-1]
}
