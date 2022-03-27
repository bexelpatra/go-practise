package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek0000() {
	r := bufio.NewReader(os.Stdin)
	var n,m int
	fmt.Fscanln(r,&n,&m)

	graph := make([][]int,n+1)
	visited := make([][]int,n+1)
	for i,_ :=range graph{
		graph[i] = make([]int, m+1)
		visited[i] = make([]int, m+1)
	}

	for i := 0; i < m; i++ {
		var a,b int
		fmt.Fscanln(r, &a,&b)
		graph[a][b] = 1
		graph[b][a] = 1
	}


}