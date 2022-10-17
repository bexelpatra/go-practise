package baek

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func Baek22173() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	t:= scanInt22173(r)
	n:= scanInt22173(r)

	// pss := make([][3]int,n)
	processQ := &ProcessQ{}
	
	for i := 0; i < n; i++ {
		heap.Push(processQ,MyProcess{scanInt22173(r),scanInt22173(r),scanInt22173(r)})
	}
	heap.Init(processQ)
	
	sb := bytes.Buffer{}
	for i := 0; i < t; i++ {
		now :=heap.Pop(processQ).(MyProcess)
		now.order-=1
		// fmt.Println(processQ)
		sb.WriteString(strconv.Itoa(now.pid))
		sb.WriteString("\n")
		now.time-=1
		if(now.time >0){
			heap.Push(processQ,now)
			// fmt.Println(processQ)
		}
		// fmt.Println()
	}

	fmt.Println(sb.String())
}

type MyProcess struct{
	pid int
	time int
	order int
}

type ProcessQ []MyProcess

func (pq ProcessQ) Len() int { return len(pq) }
func (pq ProcessQ) Less(i, j int) bool {
	if pq[i].order == pq[j].order {
		return pq[i].pid < pq[j].pid
	}else{
		return pq[i].order>pq[j].order
	}
}
func (pq ProcessQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *ProcessQ) Push(x interface{}) { 
	*pq = append(*pq, x.(MyProcess)) 
}

func (pq *ProcessQ) Pop() (popped interface{}) {
	popped = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return
}

func scanInt22173(r *bufio.Scanner)int{
	r.Scan()
	num :=0
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}