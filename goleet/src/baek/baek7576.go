package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek7576() {
	r := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscanln(r, &M, &N)

	graph := make([][]int, N)
	visited := make([][]int, N)
	q := make([]point, 0)
	for i, _ := range graph {
		temp := make([]int, M)
		visited[i] = make([]int, M)

		line, _, _ := r.ReadLine()
		rows := strings.Split(string(line), " ")

		for j, num := range rows {
			temp[j], _ = strconv.Atoi(num)

			if num == "1" {
				q = append(q, point{i, j})
				visited[i][j] = 1
			}
		}
		graph[i] = temp
	}

	for len(q) > 0 {
		x := q[0].x
		y := q[0].y
		q = q[1:]

		bfs7576(&q, &graph, &visited, x, y, x-1, y, N, M)
		bfs7576(&q, &graph, &visited, x, y, x+1, y, N, M)
		bfs7576(&q, &graph, &visited, x, y, x, y-1, N, M)
		bfs7576(&q, &graph, &visited, x, y, x, y+1, N, M)

	}
	max := 0

	for i, _ := range graph {
		for _, num := range graph[i] {
			if num == 0 {
				fmt.Println("-1")
				return
			}
			max = func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}(max, num)
		}
	}
	fmt.Println(max - 1)

}

func bfs7576(q *[]point, graph *[][]int, visited *[][]int, a, b, x, y, N, M int) {
	if x < 0 || y < 0 || x >= N || y >= M {
		return
	}
	if (*visited)[x][y] == 1 || (*graph)[x][y] == -1 {
		return
	}

	if (*graph)[x][y] == 0 {
		(*graph)[x][y] = (*graph)[a][b] + 1
		(*visited)[x][y] = 1
		*q = append(*q, point{x, y})
		return
	}
}

type point struct {
	x, y int
}
