package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_strings(t *testing.T) {
	a := "abcdefgh"

	for _, N := range a {
		fmt.Println(N, string(N))
	}
	num, err := strconv.ParseInt("1012", 3, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(num)
}

func Test_slice(t *testing.T) {
	q := make([]int, 0)
	q = append(q, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	q = append(q[:2], q[3:]...)
	fmt.Println(q)

}

func Test_mapset(t *testing.T) {

}
