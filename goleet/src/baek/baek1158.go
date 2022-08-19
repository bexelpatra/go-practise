package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1158() {
	// 길이가 길수록 시간이 오래 걸린다.
	// 리스트에서 내용을 제거하지 않고있기때문
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, k int
	fmt.Fscanln(r, &n, &k)
	circle := make([]int, n)
	answer := make([]int, 0)
	for i := 0; i < n; i++ {
		circle[i] = i + 1
	}
	now := 0
	next := -1
	for len(answer) < n {
		for now < k {
			next = (next + 1) % n
			if circle[next] != 0 {
				now += 1
			}
		}
		answer = append(answer, circle[next])
		circle[next] = 0
		now = 0
	}
	fmt.Fprintf(w, "<")
	for i := range circle {

		fmt.Fprint(w, answer[i])
		if i == len(circle)-1 {
			break
		}
		fmt.Fprint(w, ", ")
	}
	fmt.Fprintf(w, ">")

}

func Baek1158_2() {
	// r := bufio.NewReader(os.Stdin)
	// w := bufio.NewWriter(os.Stdout)
	// defer w.Flush()
	// var n, k int
	// fmt.Fscanln(r, &n, &k)
	// circle := make([]int, n)
	// answer := make([]int, 0)
	// for i := 0; i < n; i++ {
	// 	circle[i] = i + 1
	// }

}

func Baek1158_3() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, k int
	fmt.Fscanln(r, &n, &k)
	circle := make([]int, n)
	answer := make([]int, 0)
	for i := 0; i < n; i++ {
		circle[i] = i + 1
	}

}