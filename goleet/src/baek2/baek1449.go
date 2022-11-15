package baek2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek1449() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:= scanInt1449(r)
	k:= scanInt1449(r)

	pipe :=make([]bool,1050)
	
	for i := 0; i < n; i++ {
		n :=scanInt1449(r)-1
		pipe[n] = true
	}
	result :=0
	
	for idx := 0;idx < 1000; idx++{
		if pipe[idx]{
			result+=1
			for i := idx; i < idx+k; i++ {
				pipe[i] = false
			}
			idx+=k-1
			if idx>=1000{
				break
			}
		}
	}
	fmt.Println(result)
}
func Baek1449_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:= scanInt1449(r)
	k:= scanInt1449(r)

	pipe :=make([]int,n)
	
	for i := 0; i < n; i++ {
		pipe[n] = scanInt1449(r)
	}
	sort.Ints(pipe)
	result :=0
	position :=0
	for i := 0; i < n; i++ {
		if pipe[i]> position{
			result+=1
			position+= pipe[i]+k-1
		}
	}
	
	fmt.Println(result)
}

func scanInt1449(r *bufio.Scanner) int {
	num:=0
	r.Scan()
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')	
	}
	return num
}
