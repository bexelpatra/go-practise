package baek2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek1399() {
	r := bufio.NewReader(os.Stdin)
	n,_,_:=r.ReadLine()
	N:=0
	for _, v := range n {
		N*=10
		N+=(int)(v-'0')
	}
	alphabet :=make([]int,26)
	for i := 0; i < N; i++ {
		text,_,_:=r.ReadLine()
		for j := 0; j < len(text); j++ {
			alphabet[(int)(text[j]-'A')]+= pow1399(1,len(text)-j)
		}
	}
	sort.Slice(alphabet,func(i, j int) bool {return alphabet[i]>alphabet[j]})
	power :=9
	result :=0
	for _, v := range alphabet {
		result+= (v*power)
		power--
	}
	fmt.Println(result)
}

func pow1399(n,pow int) int{
	for i := 1; i < pow; i++ {
		n*=10	
	}
	return n
}
