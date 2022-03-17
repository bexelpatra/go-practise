package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1012() {
	r := bufio.NewReader(os.Stdin)
	var T int
	answer := make([]int, 0)

	fmt.Fscanln(r, &T)
	for i := 0; i < T; i++ {
		var X, Y, K int
		var graph [][]int
		var checked [][]int
		fmt.Fscanln(r, &X, &Y, &K)

		graph = make([][]int, X+1)
		checked = make([][]int, X+1)
		for i := 0; i < X+1; i++ {
			graph[i] = make([]int, Y+1)
			checked[i] = make([]int, Y+1)
		}

		for i := 0; i < K; i++ {
			var x, y int
			fmt.Fscanln(r, &x, &y)
			graph[x][y] = 1
		}
		total := 0
		for i := 0; i < X; i++ {
			num := 0
			for j := 0; j < Y; j++ {
				dfs1012(graph, checked, i, j, X, Y, &num)
				if num > 0 {
					// fmt.Println(num)
					total += 1
					num = 0
				}
			}
		}
		answer = append(answer, total)
	}

	for i := range answer {
		fmt.Println(answer[i])
	}
}

func dfs1012(graph, checked [][]int, x, y, X, Y int, cnt *int) {
	if x < 0 || y < 0 || x >= X || y >= Y {
		return
	}
	if checked[x][y] == 1 {
		return
	}
	if graph[x][y] == 0 {
		return
	}
	checked[x][y] = 1
	*cnt += 1
	if graph[x][y] == 1 {
		dfs1012(graph, checked, x-1, y, X, Y, cnt)
		dfs1012(graph, checked, x+1, y, X, Y, cnt)
		dfs1012(graph, checked, x, y-1, X, Y, cnt)
		dfs1012(graph, checked, x, y+1, X, Y, cnt)
	}
}
