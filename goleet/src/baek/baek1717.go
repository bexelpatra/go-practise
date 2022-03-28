package baek

import (
	"bufio"
	"fmt"
	"os"
)

const (
	yes = "YES"
	no  = "NO"
)

func Baek1717() {
	// 실패
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var n, m int
	fmt.Fscanln(r, &n, &m)
	list := make([][]int, n+1)
	set := make([]map[int]struct{}, n+1)
	for i := range list {
		list[i] = make([]int, 0)
		list[i] = append(list[i], i)

		set[i] = make(map[int]struct{})
		set[i][i] = struct{}{}
	}

	for i := 0; i < m; i++ {
		var op, a, b int
		fmt.Fscanln(r, &op, &a, &b)

		if op == 0 {
			sum(set, a, b)
			// for i := range set {
			// 	fmt.Println(set[i])
			// }
		} else {
			answer := no
			if contain(set[a], b) {
				answer = yes
			}
			fmt.Fprintln(w, answer)
		}
	}
}

func Baek1717_2() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var n, m int
	fmt.Fscanln(r, &n, &m)
	list := make([]int, n+1)
	for i := range list {
		list[i] = i
	}
	for i := 0; i < m; i++ {
		var op, a, b int
		fmt.Fscanln(r, &op, &a, &b)

		if op == 0 {
			unionParent(list, a, b)
		} else {
			answer := no
			if findParent(list, a, b) == 1 {
				answer = yes
			}
			fmt.Fprintln(w, answer)
		}
	}
}

func sum(list []map[int]struct{}, a, b int) []map[int]struct{} {
	for key, _ := range list[a] {
		list[key], list[b] = concat(list[key], list[b])
	}
	for key, _ := range list[b] {
		list[key], list[a] = concat(list[key], list[a])
	}
	return list
}

func concat(a, b map[int]struct{}) (x, y map[int]struct{}) {
	for k, v := range a {
		b[k] = v
	}
	for k, v := range b {
		a[k] = v
	}
	return a, b
}

func contain(m map[int]struct{}, i int) bool {
	_, bool := m[i]
	return bool
}

func getParent(parent []int, x int) int {
	if parent[x] == x {
		return x
	}
	parent[x] = getParent(parent, parent[x])
	return parent[x]
}
func unionParent(parent []int, a, b int) []int {
	a = getParent(parent, a)
	b = getParent(parent, b)
	if a < b {
		parent[b] = a
	} else {
		parent[a] = b
	}
	return parent
}

func findParent(parent []int, a, b int) int {
	a = getParent(parent, a)
	b = getParent(parent, b)

	if a == b {
		return 1
	} else {
		return 0
	}
}
