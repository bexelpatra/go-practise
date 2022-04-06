package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 스티커
func Baek9465() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r.Split(bufio.ScanWords)
	loop := scanInt9465(r)

	for i := 0; i < loop; i++ {
		solution9465(r, w)
	}
}

func solution9465(r *bufio.Scanner, w *bufio.Writer) {

	n := scanInt9465(r)
	graph := make([][]int, 2)
	top := make([]int, n+1)
	bottom := make([]int, n+1)

	for i := 0; i <= 1; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			graph[i][j] = scanInt9465(r)
		}
	}
	top[1] = graph[0][1]
	bottom[1] = graph[1][1]

	for i := 2; i < n+1; i++ {
		top[i] = max(bottom[i-1], bottom[i-2]) + graph[0][i]
		bottom[i] = max(top[i-1], top[i-2]) + graph[1][i]
	}
	fmt.Println(top)
	fmt.Println(bottom)
	fmt.Fprintln(w, max(top[n], bottom[n]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func scanInt9465(r *bufio.Scanner) int {
	r.Scan()
	temp, _ := strconv.Atoi(r.Text())
	return temp
}
