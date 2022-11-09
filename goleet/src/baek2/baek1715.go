package baek2

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func Baek1715() {
	r:= bufio.NewScanner(os.Stdin)
	N:= scanInt1715(r)
	list := make([]int64,0)

	if N==1{
		fmt.Println(scanInt1715(r))
		return
	}
	for i := 0; i < int(N); i++ {
		list = append(list, scanInt1715(r))
	}
	sort.Slice(list,func(i, j int) bool {
		return list[i]<list[j]
	})
	var result int64
	for i := 1; i < int(N); i++ {
		list[i] += list[i-1]
		sorting1715(&list)
		result+=list[i]
	}
	fmt.Println(result)
}

func Baek1715_2(){
	r:= bufio.NewScanner(os.Stdin)
	N:= scanInt1715(r)
	list := new(PQ)
	heap.Init(list)

	for i := 0; i < int(N); i++ {
		heap.Push(list,scanInt1715_2(r))
	}
	var result int64
	for list.Len()>1{
		a:= heap.Pop(list).(int64)
		b:= heap.Pop(list).(int64)
		result+= (int64)(a+b)
		heap.Push(list,a+b)
	}
	fmt.Println(result)
	
}
type PQ []int
func (pq PQ)Len()int {return len(pq)}
func (pq *PQ)Push(a interface{}){
	(*pq) = append((*pq), a.(int))
}
func (pq *PQ) Swap(a,b int){(*pq)[a],(*pq)[b]=(*pq)[b],(*pq)[a]}
func (pq *PQ) Pop() (interface{}) {
	l :=len(*pq)
	q := (*pq)[l-1]
	(*pq)= (*pq)[:l-1]
	return q
}
func (pq PQ)Less(a,b int)bool {
	return pq[a]<pq[b]
}


func scanInt1715(r *bufio.Scanner)int64{
	var num int64
	r.Scan()
	for _, v := range r.Bytes() {
		num*=10
		num+=(int64)(v-'0')
	}
	return num
}
func scanInt1715_2(r *bufio.Scanner)int{
	var num int
	r.Scan()
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}
func sorting1715(list *[]int64){
	sort.Slice(*list,func(i, j int) bool {
		return (*list)[i]<(*list)[j]
	})
}
