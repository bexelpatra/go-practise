package main

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"

	"example.go/goleet/util"
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
	m := make(map[string]int)
	m["key1"] = 1
	delete(m, "key1")
	fmt.Println(m)

	var exits = struct{}{}
	var a interface{}
	exits = struct{}{}
	fmt.Println(exits, a)
}

func Test_emptyStructure(t *testing.T) {

	// car := struct{
	//           Speed  int
	//           Weight float
	//        }
	// you define a struct be do now create an instance and assign it to car
	// the following however is completely valid

	// car2 := struct {
	// 	Speed  int
	// 	Weight float32
	// }{6, 7.1}

	//car2 now has a Speed of 6 and Weight of 7.1
	n := make(map[string]struct{})
	m := make(map[string]struct {
		x int
		y int
	})
	k := map[string]int{}
	l := map[string]struct{}{}

	m["a"] = struct {
		x int
		y int
	}{1, 2}

	k["b"] = 1
	m["c"] = struct {
		x int
		y int
	}{}
	n["c"] = struct{}{}
	var xx string
	l["a"] = struct{}{}
	fmt.Println(xx)
	fmt.Println(xx == "")
	fmt.Println(unsafe.Sizeof(xx))
	fmt.Println(n["c"] == struct{}{})
}

func Test_se(t *testing.T) {
	s := util.NewSet()
	s.Add("a")

	s.Remove("a")
	fmt.Println(s.Contains("a"))
}
