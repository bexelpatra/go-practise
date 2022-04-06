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
	fmt.Println(loop)

	for i := 0; i < loop; i++ {
		size := scanInt9465(r)
		graph := make([][]int, 2)
		for j := 0; j < 2; j++ {
			graph[j] = make([]int, size)
			for k := range graph {
				graph[k] = scanInt11650(r)
			}
		}
		fmt.Println(graph)
	}
}

func scanInt9465(r *bufio.Scanner) int {
	r.Scan()
	temp, _ := strconv.Atoi(r.Text())
	return temp
}
