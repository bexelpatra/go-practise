package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	a := A{"1", "2", "3", "aaaa", 1000}
	strList := []string{"4", "5", "6", "newAz"}
	ayo := A{}
	myA := map[string]interface{}{}

	data, _ := json.Marshal(a)
	json.Unmarshal(data, &myA)
	fmt.Println(myA)

	s := reflect.ValueOf(&ayo).Elem()
	for i := range strList {
		f := s.Field(i)
		f.SetString(strList[i])
		fmt.Println(f.Elem())
	}

	fmt.Println(ayo)
}

type A struct {
	A string
	B string
	C string
	D string
	F int64
}

func Test_SplitStr(t *testing.T) {
	str := ""
	str1 := "1,2,3,4,"
	list := strings.Split(str, "")
	list11 := strings.Split(str, ";")

	list2 := strings.Split(str1, ",")
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(list11)
	for i := range list2 {
		fmt.Println(i, list2[i])
	}

	waet := ""

	fmt.Println(strings.LastIndex(waet, "C"))
	fmt.Println(strings.Index(waet, "C"))

}

func Test_Unmarshal(t *testing.T) {
	var a int64
	a = 1000
	b := A{}
	b.F = 1000
	fmt.Println(a)
	fmt.Printf("%#v\n", b)

	str := "100000.00"
	str2 := "10000000"

	fmt.Println(strings.Contains(str, "."))
	fmt.Println(strings.Index(str, "."))

	fmt.Println(strings.Contains(str2, "."))
	fmt.Println(strings.Index(str2, "."))
}

func Test_Spew(t *testing.T) {
	expectations := map[string]*TVNameParseResult{
		"Sleepy.Hollow.S03E01.720p.HDTV.x264-AVS": {
			Name:        "Sleepy Hollow",
			CleanedName: "Sleepy Hollow S03E01 720p HDTV x264-AVS",
			Country:     "",
			Episode:     "1",
			Season:      "3",
			Airdate:     "",
		},
		"NBC.Nightly.News.2016.02.17.WEB-DL.x264-2Maverick": {
			Name:        "NBC Nightly News",
			CleanedName: "NBC Nightly News 2016 02 17 WEB-DL x264-2Maverick",
			Country:     "",
			Episode:     "02/17",
			Season:      "2016",
			Airdate:     "2016.02.17",
		},
	}
	for name, v := range expectations {
		t.Errorf("Diff in Parse output for %s", name)
		fmt.Println("Expected:")
		spew.Dump(v, "a")
		fmt.Println("sdump")

	}
}

type TVNameParseResult struct {
	Name        string
	CleanedName string
	Country     string
	Episode     string
	Season      string
	Airdate     string
}

/*

젠킨스 빌드
개발에서 패키지
웹 콘솔에서 채널에 체인코드 업로드
커밋
승인

*/

func Test_mytest(t *testing.T) {
	bytes := []byte("23")
	minus := false
	num := 0
	for _, b := range bytes {
		num *= 10
		if b-'0' > '9' {
			minus = true
		} else {
			num += int(b - '0')
		}
	}
	if minus {
		num = -num
	}
	fmt.Println(num)
}

func Test_SliceReference(t *testing.T) {

}

func refer(dp *[][]int, a int) {
	fmt.Println((*dp)[a][0])
}

func Test_stringIndexing(t *testing.T){
	// 1. 문자열 더하기? 
	fmt.Println("123"+"456")

	s:=strings.TrimRight("1234   "," ")
	fmt.Println(s , len(s))

	s2:= strings.Builder{}
	s2.WriteByte(byte('3'))
	s2.Write([]byte("123"))
	
	fmt.Println(s2.String())
}

func TestString(t *testing.T) {
// 1.
  str1 := "Welcome "
  str2 := "Rain!"
  str := str1 + str2
  fmt.Println(str)
  str = str1 + "my " + str2
  fmt.Println(str)
  // 2.
  var b bytes.Buffer
  b.WriteString("R")
  b.WriteString("a")
  b.WriteString("i")
  b.WriteString("n")
  fmt.Println(b.String())
  // 3.
  str = fmt.Sprintf("%s%s", str1, str2)
  fmt.Println(str)
  // 4.
  str = ""
  str += str1
  str += str2
  fmt.Println(str)
  // 5.
  mySlice := []string{"Welcome", "my", "Rain!"}
  str = strings.Join(mySlice, " * ")
  fmt.Println(str)
}