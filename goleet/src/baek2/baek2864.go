package baek2

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2864() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	minA,maxA := scanInt2864(r)
	minB,maxB := scanInt2864(r)

	fmt.Println(minA+minB,maxA+maxB)
}

func scanInt2864(r *bufio.Scanner) (int,int) {
	min :=0
	max :=0
	r.Scan()
	for _,v :=range r.Bytes(){
		min*=10
		max*=10
		if (v-'0')==5 || (v-'0')==6{
			min+=5
			max+=6
		}else{
			min+=(int)(v-'0')
			max+=(int)(v-'0')
		}
	}	
	return min,max
}	