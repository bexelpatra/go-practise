package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek10610() {
	r:=bufio.NewScanner(os.Stdin)
	numbers := scanInt10610(r)
	
	stop := true
	for _, v := range numbers {
		if(v==0){
			stop = false
			break
		}
	}
	if(stop){
		fmt.Println(-1)
		return
	}	
	
}

func scanInt10610(r *bufio.Scanner) ([]int){
	r.Scan()
	bytes := r.Bytes()
	length := len(bytes)
	numbers := make([]int,length)

	for i := 0; i < length; i++ {
		numbers[i] = int(bytes[i]-'0')
	}

	return numbers
}