package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1789() {
	r:=bufio.NewScanner(os.Stdin)
	var n int64
	scanInt1789(r,&n)

	var sum int64
	sum =0
	var i int64
	if(n<3){
		fmt.Println(1)
		return
	}
	for i = 1; i < n; i++ {
		if(sum+i >=n){
			if ((sum+i)==n){
				fmt.Println(i)
				return 
			}
			fmt.Println(i-1)
			return 
		}
		sum+=i
	}
}

func scanInt1789(r *bufio.Scanner,num *int64){
	r.Scan()
	for _, v := range r.Bytes() {
		*num*=10
		*num+=(int64)(v-'0')
	}
}