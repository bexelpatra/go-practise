package baek2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek1744() {
	r:= bufio.NewScanner(os.Stdin)
	minus :=make([]int,0)
	plus := make([]int,0)
	zero :=0
	n:=scanInt1744(r)
	for i := 0; i < n; i++ {
		now := scanInt1744(r)
		if now>0{
			plus = append(plus, now)
		}else if now ==0{
			zero+=1
		}else{
			minus = append(minus, now)
		}
	}
	sort.Slice(minus,func(i, j int) bool {return minus[i]<minus[j]})
	sort.Slice(plus,func(i, j int) bool {return plus[i]<plus[j]})
	sum :=0
	if len(minus)%2==0 || len(minus)==0{ // 모두 +

	}else{ // 0의 유무를 확인해야함
		if zero>0{
			minus = minus[:len(minus)-1]
		}else{	
			sum += minus[len(minus)-1]
			minus = minus[:len(minus)-1]
		}
	}
	for i := 0; i < len(minus)-1; i+=2 {
		sum+=(minus[i]*minus[i+1])
	}

	if len(plus)%2==0|| len(plus)==0{

	}else{
		sum+= plus[0]
		plus = plus[1:]
	}

	for i := 0; i < len(plus)-1; i+=2 {
		if plus[i]*plus[i+1] > plus[i+1]{
			sum+=(plus[i]*plus[i+1])
		}else{
			sum+=(plus[i]+plus[i+1])
		}
	}
	fmt.Println(sum)
}

func scanInt1744(r *bufio.Scanner)(int){
	num:=0
	r.Scan()
	b := r.Bytes()
	minus := false	
	if b[0]=='-'{
		minus=true
	}else{
		num+=(int)(b[0]-'0')
	}
	for i := 1; i < len(b); i++ {
		num*=10
		num+=(int)(b[i]-'0')
	}
	if minus{
		num*=-1
	}
	return num
}