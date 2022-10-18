package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var (
	V   int
	E   int
	MAX int
)

func Baek1753() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	V = scanInt1753(r)
	E = scanInt1753(r)
	root := scanInt1753(r)
	MAX = 10_000_000
	graph := make([][]int, V+1)

	for i := range graph {
		graph[i] = make([]int, V+1)
	}
	for i := 0; i < E; i++ {
		a := scanInt1753(r)
		b := scanInt1753(r)
		c := scanInt1753(r)
		if graph[a][b] == 0 {
			graph[a][b] = c
		} else {

			graph[a][b] = func(x, y int) int {
				if x > y {
					return y
				}
				return x
			}(c, graph[a][b])
		}

	}

	dp := make([]int, V+1)
	for v := range dp {
		dp[v] = MAX
	}
	recursive1753(&graph, &dp, root, 0)

	result := bytes.Buffer{}
	result.WriteString("0\n")
	for i := 2; i < len(dp); i++ {
		v := (dp)[i]
		if v == MAX {
			result.WriteString("INF\n")
		} else {
			result.WriteString(strconv.Itoa(v))
			result.WriteString("\n")

		}
	}

	fmt.Println(result.String())
}

func recursive1753(graph *[][]int, dp *[]int, now int, price int) {
	for i := now + 1; i < len((*graph)[now]); i++ {
		v := (*graph)[now][i]
		if v > 0 {
			(*dp)[i] = func(a, b int) int {
				if a > b {
					return b
				}
				return a
			}(v+price, (*dp)[i])
			recursive1753(graph, dp, i, (*dp)[i])
		}
	}
}

func scanInt1753(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
