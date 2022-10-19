package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var root int

// map을 이용한 방식
func Baek1753_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	V := scanInt1753_2(r)
	E := scanInt1753_2(r)
	root = scanInt1753_2(r)
	graph := make(map[int]map[int]int, V+1)

	for i := 0; i < E; i++ {
		a := scanInt1753_2(r)
		b := scanInt1753_2(r)
		c := scanInt1753_2(r)

		if graph[a] == nil {
			graph[a] = make(map[int]int)
		}
		if graph[a][b] == 0 {
			graph[a][b] = c
		} else {
			graph[a][b] = func(a, b int) int {
				if a > b {
					return b
				}
				return a
			}(c, graph[a][b])
		}
	}

	dp := make(map[int]int)

	recursive1753_2(&graph, &dp, root, 0)

	result := bytes.Buffer{}
	result.WriteString("0\n")

	for i := 2; i < V+1; i++ {
		if dp[i] == 0 {
			result.WriteString("INF\n")
		} else {
			result.WriteString(strconv.Itoa(dp[i]))
		}
		result.WriteString("\n")
	}
	fmt.Println(result.String())
}

func recursive1753_2(graph *map[int]map[int]int, dp *map[int]int, now int, price int) {
	for key, value := range (*graph)[now] {
		if (*dp)[key] == 0 {
			(*dp)[key] = value + price
		} else {
			(*dp)[key] = func(a, b int) int {
				if a > b {
					return b
				}
				return a
			}((*dp)[key], value+price)
		}
		if key != root {
			recursive1753_2(graph, dp, key, price+value)
		}
	}
}

func scanInt1753_2(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
