package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

type Sample_1 struct {
	A  string
	B  string
	Ab string
}
type Sample_2 struct {
	A string
	B string
}

func Test_Marshal(t *testing.T) {
	a := Sample_1{"a", "b", "c"}
	a2 := Sample_2{}
	result, _ := json.Marshal(a)
	fmt.Println(string(result))

	err := json.Unmarshal(result, &a2)
	if err != nil {
		panic("c8")
	}

	fmt.Println(print(json.Marshal(a2)))

}

func print(data []byte, a interface{}) string {
	return string(data)
}

func Test_LastIndex(t *testing.T) {
	str := "abcdefc"
	fmt.Println(strings.LastIndex(str, "c"))
}

func Test_DeleteMap(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "a"
	m["b"] = "b"
	m["c"] = "c"

	fmt.Println(m)
	delete(m, "d")
	fmt.Println(m)

}

func Test_stringcompare(t *testing.T) {

	a := "abc"
	a2 := "abc"
	b := "bcd"
	b2 := "abcc"
	fmt.Println(a < b)
	fmt.Println(a == a2)
	fmt.Println(b < b2)

	fmt.Println("" < "20222020")
	fmt.Println("" < "2")
}

func Test_Permutation(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := counter()
	perm_test1(list, 3, 0, c)
	fmt.Println(c() - 1)
}
func counter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}

func perm_test1(list []int, n int, r int, c func() int) {

	if n == r {
		fmt.Println(list[:n])
	}
	c()
	for i := r; i < len(list); i++ {
		list[r], list[i] = list[i], list[r]
		perm_test1(list, n, r+1, c)
		list[r], list[i] = list[i], list[r]
	}
}

var myLogger *log.Logger

func Test_time(t *testing.T) {
	now := time.Now()
	custom := now.Format("20060102")
	txt := fmt.Sprintf("logfile_%s.txt", custom)

	fpLog, err := os.OpenFile(txt, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	myLogger.Println("End of Program")

}
func Test_defer(t *testing.T) {
	// return 치고 마지막에 defer 함수 실행
	fmt.Println(sample_defer())

}

func sample_defer() string {
	str := "123"
	defer func() {
		str = "456"
		fmt.Println("안 바뀜")
	}()
	return str
}
