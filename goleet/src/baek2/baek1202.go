package baek2

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func Baek1202() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	N:= scanInt1202(r)
	K:= scanInt1202(r)
	jewerlies := make([]Jewerly1202,N)
	C:=make([]int,K)
	for i := 0; i < N; i++ {
		jewerlies[i] = Jewerly1202{scanInt1202(r),scanInt1202(r)}
	}
	for i := 0; i < K; i++ {
		C[i] = scanInt1202(r)
	}
	sort.Slice(jewerlies,func(i, j int) bool {
		if jewerlies[i].M == jewerlies[j].M{
			return jewerlies[i].V > jewerlies[j].V
		}else{
			return jewerlies[i].M < jewerlies[j].M
		}
	})
	sort.Slice(C,func(i, j int) bool {return C[i]<C[j]})
	
 	var result int64
	var idx int
	for _, c := range C {
		
		for idx = 0; idx < N; idx++ {
			if jewerlies[idx].M <= c {
				result+= (int64)(jewerlies[idx].V)
				// jewerlies[idx].status=true
				break
			}
			if jewerlies[idx].M > c{
				break
			}
		}
	}
	// fmt.Println(jewerlies)
	// fmt.Println(C)
	fmt.Println(result)

}

type Jewerly1202 struct{
	M int
	V int
	// status bool
}
type PQ1202 []int
func (pq PQ1202)Len()int {return len(pq)}
func (pq *PQ1202)Push(a interface{}){
	(*pq) = append((*pq), a.(int))
}
func (pq *PQ1202) Swap(a,b int){(*pq)[a],(*pq)[b]=(*pq)[b],(*pq)[a]}
func (pq *PQ1202) Pop() (interface{}) {
	l :=len(*pq)
	q := (*pq)[l-1]
	(*pq)= (*pq)[:l-1]
	return q
}
func (pq PQ1202)Less(a,b int)bool {
	return pq[a]<pq[b]
}

func  Pop1202(i int,jew *[]Jewerly1202){
	before := (*jew)[:i]
	after := make([]Jewerly1202,0)
	if len(*jew)-1<i{
		after = (*jew)[i+1:]
	}
	*jew = append(before,after...)
}

func scanInt1202(r *bufio.Scanner) int {
	num :=0
	r.Scan()
	for _,v :=range r.Bytes(){
		num *=10
		num+=(int)(v-'0')
	}
	return num
}


func Baek1202_2(){
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	N:= scanInt1202(r)
	K:= scanInt1202(r)
	jewerlies := make([]Jewerly1202,N)
	C:=make([]int,K)
	for i := 0; i < N; i++ {
		jewerlies[i] = Jewerly1202{scanInt1202(r),scanInt1202(r)}
	}
	for i := 0; i < K; i++ {
		C[i] = scanInt1202(r)
	}
	sort.Slice(jewerlies,func(i, j int) bool {
		return jewerlies[i].M < jewerlies[j].M
	})
	sort.Slice(C,func(i, j int) bool {return C[i]<C[j]})
	
 	var result int64
	var idx int
	tempArr := PQ11000{}
	for _, c := range C {
		for idx < N && jewerlies[idx].M <= c{
			heap.Push(&tempArr,jewerlies[idx].V)
		}
		
		if len(tempArr) !=0{
			result += (heap.Pop(&tempArr).(int64))
		}
	}
	// fmt.Println(jewerlies)
	// fmt.Println(C)
	fmt.Println(result)

}

func Baek1202_3(){
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:= scanInt1202(r)
	m:= scanInt1202(r)
	jewerlies := make([]Jewerly1202,n)
	bag:=make([]int,m)
	for i := 0; i < n; i++ {
		jewerlies[i] = Jewerly1202{scanInt1202(r),scanInt1202(r)}
	}
	for i := 0; i < m; i++ {
		bag[i] = scanInt1202(r)
	}
	sort.Slice(bag,func(i, j int) bool {return bag[i]<bag[j]})
	sort.Slice(jewerlies,func(i, j int) bool {
		return jewerlies[i].M < jewerlies[j].M
	})
	
 	var result int64
	idx :=0
	var arr PQ11000
	for i:=0 ;i <m;i++{
		for idx < n && jewerlies[idx].M <= bag[i]{
			heap.Push(&arr,-jewerlies[idx].V)
			idx++
		}
		
		if len(arr) !=0{
			result += (int64)(-1*heap.Pop(&arr).(int))
		}
	}
	fmt.Println(result)

}