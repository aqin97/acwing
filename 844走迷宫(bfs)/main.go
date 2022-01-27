package main

import "fmt"

const N = 110

//input
var n, m int

//intput map
var board [N][N]int

//distance from (0,0)
var distance [N][N]int

type point struct {
	x, y int
}

var queue []point

//init map and distance
func Init() {
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
}

func bfs() int {
	//enqueue
	queue = append(queue, point{0, 0})
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}

	for len(queue) != 0 {
		//handle the head of the queue
		p := queue[0]
		//traverse the 4 directions of p
		for i := 0; i < 4; i++ {
			x, y := p.x+dx[i], p.y+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && board[x][y] == 0 && distance[x][y] == -1 {
				distance[x][y] = distance[p.x][p.y] + 1
				queue = append(queue, point{x, y})
			}
		}
		queue = queue[1:]
	}

	return distance[n-1][m-1]
}

func main() {
	fmt.Scanf("%d %d\n", &n, &m)
	Init()
	fmt.Println(bfs())
}
