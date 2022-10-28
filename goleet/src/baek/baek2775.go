package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func Baek2775() {
	r:= bufio.NewScanner(os.Stdin)
	loop:= scanInt2775(r)
	apt :=make([][]int,15)
	for v := range apt {
		apt[v] = make([]int, 15)
		apt[v][1]=1
	}
	for i := range apt[0] {
		apt[0][i] = i
	}
	coordinate := make([][2]int,loop)
	maxFloor :=0
	maxHo :=0
	for i := 0; i < loop; i++ {
		coordinate[i][0] = scanInt2775(r)
		coordinate[i][1] = scanInt2775(r)
		maxFloor=max2775(coordinate[i][0],maxFloor)
		maxHo=max2775(coordinate[i][1],maxHo)
	}

	for floor := 1; floor < maxFloor+1; floor++ {
		for ho := 2; ho < maxHo+1; ho++ {
			if(apt[floor][ho]>0){
				continue
			}
			apt[floor][ho] =  apt[floor-1][ho]+apt[floor][ho-1]
		}
	}
	sb := bytes.Buffer{}
	for _, v := range coordinate {
		sb.WriteString(strconv.Itoa(apt[v[0]][v[1]]))
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
	
}
func scanInt2775(r *bufio.Scanner) int{
	r.Scan()
	num:=0
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}
func max2775(a,b int) int{
	if a>b{
		return a
	}
	return b
}