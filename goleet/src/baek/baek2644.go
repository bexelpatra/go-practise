package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2264() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt2264(r)
	t1 := scanInt2264(r)
	t2 := scanInt2264(r)

	rel := scanInt2264(r)

	graphs := make([][]int, n+1)
	checked := make([]bool, n+1)
	for i := 0; i < rel; i++ {
		graphs[i] = make([]int, 0)
	}
	for i := 0; i < rel; i++ {
		a := scanInt2264(r)
		b := scanInt2264(r)

		graphs[a] = append(graphs[a], b)
		graphs[b] = append(graphs[b], a)

	}

	q := make([][]int, 0)
	checked[t1] = true
	q = append(q, graphs[t1])

	result := 0
	for i := 0; i < n; i++ {
		tempQ := q[0]
		q = q[1:]
		result += 1
		inputQ := make([]int, 0)
		for _, temp := range tempQ {
			if temp == t2 {
				fmt.Println(result)
				return
			}
			if !checked[temp] {
				checked[temp] = true
				inputQ = append(inputQ, graphs[temp]...)
			}
		}

		if len(inputQ) == 0 {
			fmt.Println(-1)
			return
		}
		q = append(q, inputQ)
	}
}

func scanInt2264(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
