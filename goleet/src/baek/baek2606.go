package baek

import (
	"bufio"
	"fmt"
	"os"
)

// bfs 방식으로 풀이
func Baek2606() {
	io := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	r := io.Reader
	// w := io.Writer
	var node int
	var con int
	var connection [][]int
	var checked []int

	fmt.Fscanln(r, &node)
	fmt.Fscanln(r, &con)
	connection = make([][]int, node+1)
	checked = make([]int, node+1)
	for i := 0; i < con; i++ {
		var a, b int
		fmt.Fscanln(r, &a, &b)
		connection[a] = append(connection[a], b)
		connection[b] = append(connection[b], a)
	}

	q := make([]int, 0)
	q = append(q, 1)
	checked[1] = 1
	total := 0
	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		total += 1
		for _, nodes := range connection[now] {
			if checked[nodes] == 0 {
				q = append(q, nodes)
				checked[nodes] = 1
			}
		}

	}
	fmt.Println(total - 1)
}

var cnt int

func Baek2606_2() {
	io := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	r := io.Reader
	// w := io.Writer
	var node int
	var con int
	var connection [][]int
	var checked []int

	fmt.Fscanln(r, &node)
	fmt.Fscanln(r, &con)
	connection = make([][]int, node+1)
	checked = make([]int, node+1)
	for i := 0; i < con; i++ {
		var a, b int
		fmt.Fscanln(r, &a, &b)
		connection[a] = append(connection[a], b)
		connection[b] = append(connection[b], a)
	}

	dfs2606(connection, checked, 1)
	fmt.Println(cnt - 1)
}

func dfs2606(connection [][]int, checked []int, now int) {
	if checked[now] == 1 {
		return
	}
	nodes := connection[now]
	checked[now] = 1
	cnt += 1
	for _, node := range nodes {
		dfs2606(connection, checked, node)
	}
}
