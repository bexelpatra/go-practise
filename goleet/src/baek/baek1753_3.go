package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Vertex struct {
	from  int
	to    int
	price int
}

// struct 이용
var Root int

func Baek1753_3() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	V := scanInt1753_3(r)
	E := scanInt1753_3(r)
	Root = scanInt1753_3(r)
	graph := make(map[int][]Vertex, 0)
	max := 100_000_000
	for i := 0; i < E; i++ {
		a := scanInt1753_3(r)
		b := scanInt1753_3(r)
		c := scanInt1753_3(r)
		if graph[a] == nil {
			graph[a] = make([]Vertex, 0)
		}
		graph[a] = append(graph[a], Vertex{a, b, c})
	}

	dp := make(map[int]interface{})
	for i := range dp {
		dp[i] = max
	}
	recursive1753_3(&graph, &dp, Root, 0)
	result := bytes.Buffer{}
	result.WriteString("0\n")

	for i := 2; i < V+1; i++ {
		if dp[i] == nil {
			result.WriteString("INF\n")
		} else {
			result.WriteString(strconv.Itoa(dp[i].(int)))
		}
		result.WriteString("\n")
	}
	fmt.Println(result.String())
}

func recursive1753_3(graph *map[int][]Vertex, dp *map[int]interface{}, now int, price int) {
	for _, vertex := range (*graph)[now] {
		if (*dp)[vertex.to] == nil {
			(*dp)[vertex.to] = vertex.price + price
		} else {
			(*dp)[vertex.to] = func(a, b int) int {
				if a > b {
					return b
				}
				return a
			}((*dp)[vertex.to].(int), vertex.price+price)
		}
		if vertex.to != Root {
			recursive1753_3(graph, dp, vertex.to, vertex.price+price)
		}
	}
}

func scanInt1753_3(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
