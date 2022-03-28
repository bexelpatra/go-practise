package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek11724() {
	r := bufio.NewReader(os.Stdin)
	var nodes, paths int
	fmt.Fscanln(r, &nodes, &paths)

	graph := make([][]int, nodes+1)
	visited := make([]int, nodes+1)
	for i, _ := range graph {
		graph[i] = make([]int, nodes+1)
	}

	for i := 0; i < paths; i++ {
		var a, b int
		fmt.Fscanln(r, &a, &b)
		graph[a][b] = 1
		graph[b][a] = 1
	}
	cnt := 0
	for i := 1; i < nodes+1; i++ {
		if visited[i] == 0 {
			dfs11724(graph, visited, i)
			cnt += 1
		}
	}
	fmt.Println(cnt)
}
func dfs11724(graph [][]int, visited []int, idx int) {
	visited[idx] = 1
	for i := 1; i < len(graph); i++ {
		if visited[i] == 0 && graph[idx][i] == 1 {
			dfs11724(graph, visited, i)
		}
	}
}
