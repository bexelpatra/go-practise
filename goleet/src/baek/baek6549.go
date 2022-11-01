package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek6549() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	for true {
		loop := scanInt6549(r)
		if loop == 0 {
			return
		}
		tree := make([][2]int64, 4*loop)
		arr := make([]int64, loop)
		for i := 0; i < int(loop); i++ {
			arr[i] = scanInt6549(r)
		}
		init6549(0, loop-1, 1, &tree, &arr)

		for _, v := range tree {
			fmt.Println(v)
		}
	}
}
func init6549(start, end, node int64, tree *[][2]int64, arr *[]int64) [2]int64 {
	if start == end {
		(*tree)[node][0] = (*arr)[start]
		(*tree)[node][1] = (*arr)[start]
		return (*tree)[node]
	}
	mid := (start + end) / 2
	left := init6549(start, mid, node*2, tree, arr)
	right := init6549(mid+1, end, node*2+1, tree, arr)

	h := min6549(left[1], right[1]) // 낮은 높이를 구한다.
	(*tree)[node][0] = (end - start + 1) * h
	(*tree)[node][1] = h
	return (*tree)[node]
}

func scanInt6549(r *bufio.Scanner) int64 {
	var num int64
	r.Scan()
	for _, v := range r.Bytes() {
		num *= 10
		num += (int64)(v - '0')
	}
	return num
}

func min6549(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
