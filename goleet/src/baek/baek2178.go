package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2178() {

	var N, M int
	var graph [][]int
	var checked [][]int
	var matrix []string

	q := make([][]int, 0)
	X := [...]int{1, -1, 0, 0}
	Y := [...]int{0, 0, 1, -1}

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscanln(r, &N, &M)

	graph = make([][]int, N)
	checked = make([][]int, N)
	matrix = make([]string, N)

	for i := range graph {
		graph[i] = make([]int, M)
		checked[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &matrix[i])
	}
	q = append(q, []int{0, 0})

	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		n, m := now[0], now[1]
		graph[0][0] = 1
		checked[0][0] = 1
		for i := 0; i < 4; i++ {
			nextN, nextM := n+X[i], m+Y[i]
			if nextN < 0 || nextM < 0 || nextN >= N || nextM >= M {
				continue
			}
			if matrix[nextN][nextM] == '1' && checked[nextN][nextM] == 0 {
				checked[nextN][nextM] = 1
				q = append(q, []int{nextN, nextM})
				graph[nextN][nextM] = graph[n][m] + 1
			}

		}
	}

	fmt.Fprintln(w, graph[N-1][M-1])
}

func Baek2178_2() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int
	fmt.Fscanln(reader, &N, &M)

	graph := NewGraph(N, M)

	for i := range graph.matrix {
		fmt.Fscanln(reader, &graph.matrix[i])
	}

	graph.BFS(0, 0)

	fmt.Fprintln(writer, graph.answer[N-1][M-1]+1)
}

func NewGraph(n, m int) graph {
	matrix := make([]string, n)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	answer := make([][]int, n)
	for i := range answer {
		answer[i] = make([]int, m)
	}
	dn, dm := [4]int{0, 0, -1, 1}, [4]int{1, -1, 0, 0}

	return graph{matrix: matrix, visited: visited, answer: answer,
		dn: dn, dm: dm, N: n, M: m}
}

type graph struct {
	matrix  []string
	visited [][]bool
	answer  [][]int
	dn, dm  [4]int
	N, M    int
}

func (g graph) BFS(n, m int) {
	queue := NewQueue()
	g.visited[n][m] = true
	queue.Enqueue(node{n, m})

	for !queue.isEmpty() {
		deqNode := queue.Dequeue()
		n, m := deqNode.n, deqNode.m
		for i := 0; i < 4; i++ {
			currentN, currentM := n+g.dn[i], m+g.dm[i]
			if currentN >= 0 && currentM >= 0 && currentN < g.N && currentM < g.M {
				if g.matrix[currentN][currentM] == '1' && !g.visited[currentN][currentM] {
					g.answer[currentN][currentM] = g.answer[n][m] + 1
					g.visited[currentN][currentM] = true
					queue.Enqueue(node{currentN, currentM})
				}
			}
		}
	}

}

type node struct {
	n, m int
}

func NewQueue() queue {
	return queue{}
}

type queue []node

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func (q *queue) Enqueue(node node) {
	*q = append(*q, node)
}

func (q *queue) Dequeue() node {
	if q.isEmpty() {
		return node{-1, -1}
	}
	node := (*q)[0]
	(*q) = (*q)[1:]
	return node
}
