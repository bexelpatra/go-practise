package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2264() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n :=scanInt2264(r)
	t1:=scanInt2264(r)
	t2:=scanInt2264(r)

	len := scanInt2264(r)

	graphs :=make([][]int,n+1)
	for i := 0; i < len; i++ {
		graphs[i] = make([]int, n+1)
	}
	for i := 0; i < len; i++ {
		a:=scanInt2264(r)
		b:=scanInt2264(r)
		if(graphs[a]==nil){
			graphs[a] = make([]int, n+1)
		}
		if(graphs[b]==nil){
			graphs[b] = make([]int, n+1)
		}
		graphs[a][b] = 1
		graphs[b][a] = 1
	}

	fmt.Println(t1,t2)
	fmt.Println(graphs)
}

func scanInt2264(r *bufio.Scanner) int{
	r.Scan()
	num :=0
	for _, v := range r.Bytes() {
		num *=10
		num+=(int)(v-'0')
	}
	return num
}