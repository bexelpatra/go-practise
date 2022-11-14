package baek2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek1049() {
	r:=bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:= scanInt1049(r)
	m:= scanInt1049(r)
	sixString :=make([]int,m)
	oneString :=make([]int,m)
	for i := 0; i < m; i++ {
		sixString[i] = scanInt1049(r)
		oneString[i] = scanInt1049(r)
	}
	sort.Slice(sixString,func(i, j int) bool {
		return sixString[i] < sixString[j]
	})
	sort.Slice(oneString,func(i, j int) bool {
		return oneString[i] < oneString[j]
	})
	idx :=0
	maxOne := n*oneString[0]
	for n > 0{
		if sixString[0] ==0{
			fmt.Println(0)
			return
		}
		n-= 6
		idx+=1
	}

	n+=6
	idx-=1
	oneIdx :=0
	for n>0{
		if oneString[0]==0{
			fmt.Println(0)
			return
		}
		n-=1
		oneIdx+=1
	}
	
	result :=0
	if oneIdx*oneString[0] > sixString[0]{
		result= sixString[0]*(idx+1)
	}else{
		result= sixString[0]*idx + oneString[0]*oneIdx
	}
	
	result = func(a,b int) int {
		if a<b{
			return a
		}
		return b
	}(result,maxOne)
	fmt.Println(result)
}

func scanInt1049(r *bufio.Scanner)int{
	num:=0
	r.Scan()
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}