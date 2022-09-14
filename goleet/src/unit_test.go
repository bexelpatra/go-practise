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

func Test_slislice(t *testing.T) {
	sli := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	RmSlice(&sli, 1)
	fmt.Println(&sli, len(sli))
	RmSlice(&sli, 2)
	fmt.Println(&sli, len(sli))
}
func RmSlice(s *[]int, i int) {
	for _, v := range *s {
		if v == i {
			(*s)[i] = (*s)[len(*s)-1]
			(*s) = (*s)[:len(*s)-1]
			return
		}
	}
}

func Test_string(t *testing.T) {

	fmt.Println(string('2' - '0'))
	fmt.Println(string(50 - '0'))
	fmt.Println(string('0'))
	fmt.Println(int('0'))
	fmt.Println(int('1'))
	fmt.Println(int('2'))
	fmt.Println(int('3'))
	fmt.Println(int('4'))
	fmt.Println(int('5'))
	fmt.Println(int('6'))
	fmt.Println(int('7'))
	fmt.Println(int('8'))
	fmt.Println(string([]byte{48}))

	dst := make([]byte, 0)
	strconv.AppendInt(dst, 1, 1)
	fmt.Println(string(dst))
}
