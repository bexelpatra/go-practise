package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2164() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n := 0
	for _, val := range r.Bytes() {
		n *= 10
		n += int(val - '0')
	}
	list := make([]int, n)
	for i := range list {
		list[i] = i + 1
	}
	// fmt.Println(recursive2164(&list))
	fmt.Println(Sol2164_2(n))
}

func recursive2164(iList *[]int) int {
	list := (*iList)
	if len(list) <= 2 {
		return list[len(list)-1]
	}

	list = list[1:]
	list = append(list[1:], list[0])
	return recursive2164(&list)
}

func Sol2164_2(n int) int {
	q := []int{}
	for i := 0; i < n; i++ {
		q = append(q, i+1)
	}
	for len(q) > 1 {
		q = q[1:]
		q = append(q[1:], q[0])
	}
	return q[0]
}

//https://www.acmicpc.net/source/38192902
// 큐 자료구조를 사용하는것이 아니라 원리를 적용
func Sol2164_3(n int) int {
	if n == 1 {
		return 1
	}
	list := make([]int, n, 1_000_000)
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}
	first := 0
	last := n - 1

	for first != last {
		list[first] = 0
		first += 1
		if first == last {
			return list[first]
		}
		list[last+1] = list[first]
		list[first] = 0
		last += 1
		first += 1

	}
	return 0
}
