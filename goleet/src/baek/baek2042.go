package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 세그먼트 트리
func Baek2042() {

	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	N := scanInt2042(r)
	M := scanInt2042(r)
	K := scanInt2042(r)
	arr := make([]int64, N)
	tree := make([]int64, N*4)

	for idx := range arr {
		arr[idx] = scanInt2042(r)
	}
	recursive2042(0, N-1, 1, &tree, &arr)
	var i int64
	for i = 0; i < M+K; i++ {
		if scanInt2042(r) == 1 {
			index := scanInt2042(r) - 1
			change := scanInt2042(r)
			dif := change - arr[index]
			arr[index] = change
			update2042(0, N-1, index, 1, dif, &tree)
		} else {
			start := scanInt2042(r) - 1
			end := scanInt2042(r) - 1
			fmt.Println(sum2042(0, N-1, start, end, 1, &tree))
		}
	}

}

func recursive2042(start, end, node int64, tree *[]int64, arr *[]int64) int64 {
	if start == end {
		(*tree)[node] = (*arr)[start]
		return (*tree)[node]
	}
	mid := (start + end) / 2
	(*tree)[node] = recursive2042(start, mid, node*2, tree, arr) + recursive2042(mid+1, end, node*2+1, tree, arr)
	return (*tree)[node]
}

func sum2042(start, end, left, right, node int64, tree *[]int64) int64 {
	if left > end || right < start {
		return 0
	}
	if left <= start && end <= right {
		return (*tree)[node]
	}
	mid := (start + end) / 2
	return sum2042(start, mid, left, right, node*2, tree) + sum2042(mid+1, end, left, right, node*2+1, tree)
}

func update2042(start, end, index, node, diff int64, tree *[]int64) {
	if index < start || index > end {
		return
	}
	(*tree)[node] += diff
	if start == end {
		return
	}
	mid := (start + end) / 2
	update2042(start, mid, index, node*2, diff, tree)
	update2042(mid+1, end, index, node*2+1, diff, tree)
}
func scanInt2042(r *bufio.Scanner) int64 {
	r.Scan()
	num, _ := strconv.ParseInt(r.Text(), 10, 64)
	return num
}
