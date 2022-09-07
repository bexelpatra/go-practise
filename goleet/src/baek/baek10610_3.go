package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Baek10610_3() {
	r:= bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(r,&s)

	numbers :=make([]int,10)
	sum :=0
	for _, v := range s {
		numbers[v-'0']+=1
		sum+=(int)(v-'0')
	}
	if sum %3!=0 || numbers[0]==0{
		fmt.Println(-1)
		return
	}
	var b bytes.Buffer
	for i := 9; i >= 0; i-- {
		for j := 0; j < numbers[i]; j++ {
			b.WriteByte((byte)(i+'0'))
		}
	}
	fmt.Println(b.String())

}