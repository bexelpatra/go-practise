package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func Baek10162() {
	r:= bufio.NewScanner(os.Stdin)
	var b bytes.Buffer
	t := scanInt10162(r)
	timer := []int{300,60,10}
	result := make([]int,3)
	for i, v := range timer {
		result[i] = t/v
		b.WriteString(strconv.Itoa(t/v)+" ")
		t%=v
		
	}
	if t>0{
		fmt.Println(-1)
		return
	}
	fmt.Println(b.String()[0:b.Len()-1])
}
func scanInt10162(r *bufio.Scanner)int{
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num*=10
		num+=int(v-'0')
	}
	return num
}

