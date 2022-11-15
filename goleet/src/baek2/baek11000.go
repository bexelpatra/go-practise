package baek2

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func Baek11000() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	N:= scanInt11000(r)
	list := make([][2]int,N)
	for i := 0; i < N; i++ {
		list[i][0]=scanInt11000(r)
		list[i][1]=scanInt11000(r)
	}
	sort.Slice(list,func(i, j int) bool {
		if list[i][0] == list[j][0]{
			return list[i][1]<list[j][1]
		}
		return list[i][0]< list[j][0]
	})
	pq :=make(PQ11000,0)
	heap.Push(&pq,list[0][1])
	for i := 1; i < N; i++ {
		if pq.Peek() <= list[i][0]{
			heap.Pop(&pq)
		}
		heap.Push(&pq,list[i][1])
	}

	fmt.Println(pq.Len())
}
// type item11000 struct{
// 	start, end int
// }
// type items11000 []item11000

// func (s items11000)Len()int{return len(s)}
// func (s items11000)Less(a,b int)bool{return (s)[a].start<(s)[b].start}
// func (s items11000)Swap(a,b int){s[a],s[b] = s[b],s[a]}
// func (s *items11000)Push(x interface{}){*s = append(*s, x.(item11000))}
// func (s *items11000)Pop() (popped interface{} ){
// 	popped = (*s)[len(*s)-1]
// 	*s = (*s)[:len(*s)-1]
// 	return popped
// }

type PQ11000 []int
func (pq PQ11000)Len()int {return len(pq)}
func (pq *PQ11000) Swap(a,b int){(*pq)[a],(*pq)[b]=(*pq)[b],(*pq)[a]}
func (pq PQ11000)Less(a,b int)bool {return pq[a]<pq[b]}
func (pq *PQ11000)Push(a interface{}){
	(*pq) = append((*pq), a.(int))
}
func (pq *PQ11000) Pop() (interface{}) {
	l :=len(*pq)
	q := (*pq)[l-1]
	(*pq)= (*pq)[:l-1]
	return q
}
func (pq *PQ11000)Peek()int{
	return (*pq)[0]
}
func scanInt11000(r *bufio.Scanner )int{
	num:=0
	r.Scan()
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}

func Baek11000_2(){
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	N:= scanInt11000(r)

	list := make([][2]int,N*2)
	for i := 0; i < N; i+=2 {
		list[i][0] = 1
		list[i][1] = scanInt11000(r)
		list[i+1][0] = -1
		list[i+1][1] =  scanInt11000(r)
	}

	sort.Slice(list,func(i, j int) bool {
		if list[i][1]==list[j][1]{
			return list[i][0]<list[j][0]
		}
		return list[i][1]<list[j][1]
	})

	current :=0
	max :=0
	for _, v := range list {
		current+=v[0]
		if current > max{
			max=current
		}
	}
	fmt.Println(max)
}