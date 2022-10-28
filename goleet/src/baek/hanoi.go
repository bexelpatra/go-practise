package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Hanoi() {
	r := bufio.NewScanner(os.Stdin)

	n:= nextInt1003(r)
	hanoi(1,2,3,n)

}

func hanoi(from, mid, to,num int){
	if num ==0{
		return
	}
	fmt.Printf("from : %d , mid : %d to : %d num : %d \n",from,mid,to,num)
	hanoi(from,to,mid,num-1);
	fmt.Printf("%d : %d -> %d\n",num,from,to)
	hanoi(mid,to,from,num-1);
}