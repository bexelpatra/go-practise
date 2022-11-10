package baek2

import (
	"bufio"
	"fmt"
	"os"
)

func Baek4796() {
	r:= bufio.NewScanner(os.Stdin)
	w:= bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	i:=0
	for{
		i+=1
		l:=scantInt4796(r)
		p:=scantInt4796(r)
		v:=scantInt4796(r)
		if l+p+v ==0{
			break
		}
		n:= (v/p)*l 
		if v%p<=l{
			n+=+ v%p
		}else{	
			n+=+ l
		}
		w.WriteString(ans4796(i,n))
	}
}

func scantInt4796(r *bufio.Scanner)int{
	r.Scan()
	num :=0
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}

func ans4796(i,n int)string{
	return fmt.Sprintf("Case %d: %d\n",i,n)
}