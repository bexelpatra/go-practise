package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_set(t *testing.T) {

	m := map[string]int{}
	m["key"] = 1

	a, b := m["key"]
	fmt.Println(a, b)

}

func Test_outofran(t *testing.T) {
	// list := make([]int, 0)
	m := make([]map[int]struct{}, 0)

	m = append(m, map[int]struct{}{})
	m = append(m, map[int]struct{}{})
	m = append(m, map[int]struct{}{})

	for i := range m {
		m[i][3] = struct{}{}
	}
	fmt.Println(m)
}

func TestXxx(t *testing.T) {

	r := bufio.NewReader(os.Stdin)
	var str string
	fmt.Fscanln(r)
	fmt.Fscanln(r, &str)

	// len := len(str)
	total := 0
	for _, val := range str {
		total += int(val) - '0'
	}
	fmt.Println(total)
}
