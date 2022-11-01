package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var (
	_max6549 int64
)

func Baek6549() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	w := bytes.Buffer{}
	for {
		loop := scanInt6549(r)
		if loop == 0 {
			break
		}
		tree := make([][2]int64, 4*loop)
		arr := make([]int64, loop)
		for i := 0; i < int(loop); i++ {
			arr[i] = scanInt6549(r)
		}
		init6549(0, loop-1, 1, &tree, &arr)
		w.WriteString(strconv.FormatInt(_max6549, 10))
		w.WriteString("\n")

		for _, v := range tree {
			fmt.Println(v)
		}
	}
	s := w.String()
	fmt.Print(s)
}
func Baek6549_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	w := bytes.Buffer{}
	for {
		loop := scanInt6549(r)
		if loop == 0 {
			break
		}
		tree := make([][2]int64, 4*loop)
		arr := make([]int64, loop)
		for i := 0; i < int(loop); i++ {
			arr[i] = scanInt6549(r)
		}
		init6549(0, loop-1, 1, &tree, &arr)

		w.WriteString(strconv.FormatInt(maxWidth6549(0, loop-1, loop-1, &tree, &arr), 10))
		w.WriteString("\n")

		for _, v := range tree {
			fmt.Println(v)
		}
	}
	s := w.String()
	fmt.Print(s)
}
func init6549(start, end, node int64, tree *[][2]int64, arr *[]int64) [2]int64 {
	if start == end {
		(*tree)[node][0] = (*arr)[start]
		(*tree)[node][1] = (*arr)[start]
		_max6549 = max6549(_max6549, (*tree)[node][0])
		return (*tree)[node]
	}
	mid := (start + end) / 2
	left := init6549(start, mid, node*2, tree, arr)
	right := init6549(mid+1, end, node*2+1, tree, arr)

	h := min6549(left[1], right[1]) // 낮은 높이를 구한다.
	(*tree)[node][0] = (end - start + 1) * h
	(*tree)[node][1] = h
	_max6549 = max6549(_max6549, (*tree)[node][0])
	return (*tree)[node]
}
func query6549(start, end, node, nodeLeft, nodeRight int64, tree *[][2]int64, arr *[]int64) int64 {
	if nodeRight < start || end < nodeLeft {
		return 0
	}
	if start <= nodeLeft && nodeRight <= end {
		return (*tree)[node][1]
	}

	mid := (nodeLeft + nodeRight) / 2
	leftIndex := query6549(start, end, node*2, nodeLeft, mid, tree, arr)
	rightIndex := query6549(start, end, node*2+1, mid+1, nodeRight, tree, arr)

	if leftIndex == 0 {
		return rightIndex
	} else if rightIndex == 0 {
		return leftIndex
	} else {
		temp := rightIndex
		if (*arr)[leftIndex] < (*arr)[rightIndex] {
			temp = leftIndex
		}
		return temp
	}
}

func maxWidth6549(start, end, n int64, tree *[][2]int64, arr *[]int64) int64 {
	min := query6549(start, end, 1, 0, n, tree, arr)
	maxWidth := (end - start + 1) * (*arr)[min]

	if start <= min-1 {
		temp := maxWidth6549(start, min-1, n, tree, arr)
		maxWidth = max6549(maxWidth, temp)
	}
	if end >= min+1 {
		temp := maxWidth6549(min+1, end, n, tree, arr)
		maxWidth = max6549(maxWidth, temp)
	}
	return maxWidth
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
func max6549(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
