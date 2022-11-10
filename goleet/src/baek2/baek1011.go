package baek2

import (
	"bufio"
	"os"
	"strconv"
)

func Baek1011() {
	r:= bufio.NewScanner(os.Stdin)
	w:= bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r.Split(bufio.ScanWords)
	N:= scanInt1011(r)
	var i int64
	
	for i = 0; i < N; i++ {
		a:= scanInt1011(r)
		b:= scanInt1011(r)
		target := b-a
		var j int64
		var number int64
		var cnt int64
		for j = 1; number+(j*2) < target; j++ {
			number+= j*2
			cnt+=2
		}
		if target-number >j{
			cnt+=2
		}else{
			cnt+=1
		}
		w.WriteString(strconv.FormatInt(cnt,10))
		w.WriteString("\n")
	}
}
func scanInt1011(r *bufio.Scanner)int64{
	r.Scan()
	var num int64
	for _, v := range r.Bytes() {
		num*=10
		num+=(int64)(v-'0')
	}
	return num
}