package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Baek10610_2() {
	r := bufio.NewScanner(os.Stdin)
	numbers,err := scanString10610_2(r)
	var buf bytes.Buffer
	if(err !=nil){
		fmt.Println(-1)
		return 
	}
	if numbers[0] == 0 {
		fmt.Println(-1)
		return
	}
	
	var j int64

	for i := 9; i >= 0; i-- {
		for j = 0; j < numbers[i]; j++ {
			buf.WriteByte((byte)(i + '0'))
		}
	}
	fmt.Println(buf.String())

}

func scanString10610_2(r *bufio.Scanner) ([]int64,error ){
	result := make([]int64, 10)
	r.Scan()
	var sum int
	sum=0
	for _, v := range r.Bytes() {
		result[(v-'0')] += 1
		sum+=(int)(v-'0')
	}
	if(sum%3!=0){
		return nil,fmt.Errorf("")
	}

	return result,nil
}

