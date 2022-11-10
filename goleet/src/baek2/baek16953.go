package baek2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek16953() {
	r := bufio.NewReader(os.Stdin)
	
	temp, _,_ := r.ReadLine()
	ab :=strings.Split(string(temp)," ")
	a,_:=strconv.Atoi(ab[0])
	b,_:=strconv.Atoi(ab[1])
	
	cnt :=1
	for a!=b{
		if b ==0{
			fmt.Println(-1)
			return
		}
		newB := b%10
		if newB%2==0 {
			b=b/2
		}else if newB == 1{
			b=b/10
		}else{
			fmt.Println(-1)
			return
		}
		cnt+=1
	}
	
	fmt.Println(cnt)

}

func lastNum16953(n int) int{
	return n
}