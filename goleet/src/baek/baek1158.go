package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, k int
	fmt.Fscanln(r, &n, &k)
	circle := make([]int, n, n+1)

	for i := 0; i < n; i++ {
		circle[i] = i + 1
	}
	cnt := 0
	fmt.Fprint(w, "<")
	for len(circle) > 1 {

		if cnt+1 == k {
			fmt.Fprintf(w, "%d, ", circle[0])
			circle = circle[1:]
			cnt = 0
		} else {
			circle = append(circle, circle[0])
			circle = circle[1:]
			cnt += 1

		}
	}
	fmt.Fprintf(w, "%d>", circle[0])

}

func Baek1158_climbplant39() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			return
		}
	}(writer)

	var n, k int
	_, err := fmt.Fscan(reader, &n, &k)
	if err != nil {
		return
	}

	var queue []int
	var result []string

	for i := 0; i < n; i++ {
		queue = append(queue, i+1)
	}

	index := k - 1
	for len(queue) > 0 {
		item := queue[index]
		queue = append(queue[:index], queue[index+1:]...)
		result = append(result, strconv.Itoa(item))

		if len(queue) > 0 {
			index = (index + k - 1) % len(queue)
		}
	}

	_, err2 := fmt.Fprint(writer, "<"+strings.Join(result, ", ")+">")
	if err2 != nil {
		return
	}

}
