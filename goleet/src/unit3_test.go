package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var logChannel chan string

func logSetup(logFile string) {

	// 로그 파일이 없으면, 생성한다
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		f, _ := os.Create(logFile)
		f.Close()
	}

	// 로그 채널을 만든다
	logChannel = make(chan string, 100)

	// 채널을 통한 비동기 로깅
	go func() {
		// 채널이 닫힐 때까지 메시지 받으면 로깅
		for msg := range logChannel {
			f, _ := os.OpenFile(logFile, os.O_APPEND, os.ModeAppend)
			f.WriteString(time.Now().String() + " " + msg + "\n")
			f.Close()
		}
	}()
}
func Test_AsyncLogging(t *testing.T) {
	logSetup("./logfile.txt")

	go func() {
		for i := 1; i < 20; i++ {
			n := strconv.Itoa(i)
			println(n)
			logChannel <- n
		}
	}()

	go func() {
		for i := 100; i < 120; i++ {
			logChannel <- strconv.Itoa(i)
		}
	}()

	strconv.Itoa(100)
	time.Sleep(1 * time.Second)
	close(logChannel)
}

func Test_pointer(t *testing.T) {
	p1 := &Person{"a", 20}
	p2 := Person{"a", 20}
	// p2 := Person{"b", 21}
	// p3 := &Person{"c", 22}

	fmt.Println(p1 == &p2)
	fmt.Println(*p1 == p2)
	fmt.Println(&p1)

	p1.addAge(1)
	p1.addAge2(10)

	fmt.Printf("%v\n", &p2)

	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(reflect.TypeOf(p2))
	fmt.Println(&p1)
	fmt.Println(p1)
}

func TestPointerSlice(t *testing.T) {
	s := SSS{1, 2, 3, 4, 5, 6, 7, 8}
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s)
}
func (s *SSS) Pop() {
	*s = (*s)[1:]
}

type SSS []int

type Person struct {
	Name string
	Age  int
}

func (p *Person) addAge(age int) *Person {
	fmt.Println(&p, p)
	p.Age += age
	return p
}

func (p Person) addAge2(age int) *Person {
	fmt.Println(&p, p)
	p.Age += age
	return &p
}

func Test_Fscan(t *testing.T) {
	r := bufio.NewReader(os.Stdin)
	var x int
	var y int
	fmt.Fscanln(r, &x, &y)

	fmt.Println(x, y)
}
