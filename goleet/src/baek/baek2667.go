package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Main2667() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var N int
	var matrix []string
	var checked [][]int

	fmt.Fscanln(r, &N)
	matrix = make([]string, N)
	checked = make([][]int, N)

	for i := range checked {
		checked[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		fmt.Fscanln(r, &matrix[i])
	}
	num := 0
	total := 0
	result := make([]int, 0)
	for i := 0; i < N*N; i++ {
		num = 0
		dfs(matrix, checked, i/N, i%N, N, &num)
		if num != 0 {
			total += 1
			result = append(result, num)
		}
	}
	fmt.Println(total)
	sort.Slice(result, func(a, b int) bool {
		return result[a] < result[b]
	})

	for _, count := range result {
		fmt.Println(count)
	}
}

func dfs(matrix []string, checked [][]int, x, y, N int, cnt *int) {
	if x < 0 || y < 0 || x >= N || y >= N {
		return
	}
	if checked[x][y] == 1 || matrix[x][y] == '0' {
		return
	}
	checked[x][y] = 1
	*cnt += 1

	dfs(matrix, checked, x+1, y, N, cnt)
	dfs(matrix, checked, x, y+1, N, cnt)
	dfs(matrix, checked, x-1, y, N, cnt)
	dfs(matrix, checked, x, y-1, N, cnt)

}
