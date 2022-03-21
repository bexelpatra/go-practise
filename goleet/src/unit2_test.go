package main

// 1. set
// 2. rune to int
// 3. tcp message sending
// 4. heap
// 5. go routine

import (
	"bufio"
	"container/heap"
	"fmt"
	"net"
	"os"
	"testing"
	"time"
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

func Test_runetoint(t *testing.T) {

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

func Test_gorutine(t *testing.T) {
	say("sync")
	go say("Async1")
	go say("Async2")

	go say("Async3")
	time.Sleep(time.Second * 1)
	fmt.Println("done")
}
func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func Test_sendMsg(t *testing.T) {
	connection, error := net.Dial("localhost", ":8080")
	if error != nil {
		fmt.Println(error)
	} else {
		_, error := connection.Write([]byte("ddddd"))
		if error == nil {
			fmt.Println("전송 성공")
		}
	}

}
func Test_pq(t *testing.T) {
	pq := &TaskPQ{
		{3, "Vacuum"},
		{2, "Feed cat"},
		{4, "Arrange play date"},
		{1, "Pack for trip"}}

	pq.Push(Task{4, "4b"})
	pq.Push(Task{5, "5b"})
	pq.Push(Task{6, "6b"})

	fmt.Println(pq.Pop())
	// heapify
	heap.Init(pq)

	// enqueue
	heap.Push(pq, Task{6, "Iron"})
	heap.Push(pq, Task{2, "new"})
	heap.Push(pq, Task{1, "pop a perkey"})

	for pq.Len() != 0 {
		// dequeue
		fmt.Println(heap.Pop(pq))
	}
}

type Task struct {
	priority int
	name     string
}

type TaskPQ []Task

func (self TaskPQ) Len() int { return len(self) }
func (self TaskPQ) Less(i, j int) bool {
	return self[i].priority < self[j].priority
}
func (self TaskPQ) Swap(i, j int)       { self[i], self[j] = self[j], self[i] }
func (self *TaskPQ) Push(x interface{}) { *self = append(*self, x.(Task)) }

func (self *TaskPQ) Pop() (popped interface{}) {
	popped = (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return
}
