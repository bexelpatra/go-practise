package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
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

func Test_logging(t *testing.T) {
	fpLog, err := os.OpenFile("C:/c8/logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer fpLog.Close()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovery")
		}
	}()

	// 표준로거를 파일로그로 변경
	// log.SetOutput(fpLog)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("INFO: ")
	log.Println("gogogogogo")
	log.Println("End of Program")
}

func Test_Trim(t *testing.T) {
	str := "abc:123;def:456;"
	fmt.Println(strings.TrimSuffix(str, ";"))
	fmt.Println(string(rune(352)))

	fmt.Printf("\n\n\n")
	fmt.Println("d")

	fmt.Println("20221010" < "")
}

func Test_Parse(t *testing.T) {
	now := "20220401"
	compare1 := "20220402"
	compare2 := "20220331"

	fmt.Println(now > compare1)
	fmt.Println(now > compare2)
	fmt.Println(now == compare1)
	fmt.Println(now == compare2)

	list := []string{"v", "c", "t"}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	fmt.Println(list)

	fmt.Println(calc(10, 1))
	fmt.Println(calc(1, 2))
	fmt.Println(calc(5, 100))
	fmt.Println(calc(20, 14))
	fmt.Println(calc(10, 10))
}

func calc(a, b int) (int, int) {
	if a > b {
		return a - b, 0
	} else {
		return 0, b - a
	}

}

func Test_PointingList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, val := range list {
		list[i] = val + 1
	}

	fmt.Println(list)

}

func Test_PrintFormat(t *testing.T) {
	a := A{1, 2, 3, "aaaa"}

	myA := map[string]interface{}{}

	data, _ := json.Marshal(a)
	json.Unmarshal(data, &myA)
	fmt.Println(myA)
	for k, v := range myA {
		fmt.Println(k, ":", v)
	}

	fmt.Println("20220202" < "")
	fmt.Println("20220202" > "")

}

type A struct {
	A int
	B int
	C int
	D string
}
