package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek13305() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	n :=scanInt13305(r)
	distances :=make([]int, n-1)
	prices :=make([]int,n)
	for i := 0; i < len(distances); i++ {
		distances[i] = scanInt13305(r)
	}
	for i := 0; i < len(prices); i++ {
		prices[i] = scanInt13305(r)
	}
	min :=1_000_000_001
	result :=0
	for i := 0; i < n-1; i++ {
		min = func(a,b int)int {
			if a<b{return a}
			return b
		}(prices[i],min)
		result+= min*distances[i]
	}
	fmt.Println(result)
}

func scanInt13305(r *bufio.Scanner) int {
	r.Scan()
	num :=0
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}