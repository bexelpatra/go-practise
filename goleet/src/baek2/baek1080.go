package baek2

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1080() {
	r:=bufio.NewReader(os.Stdin)
	var y,x int
	fmt.Fscanln(r,&y,&x)
	A :=make([][]byte,y)
	B :=make([][]byte,y)
	for i := 0; i < y; i++ {
		A[i] = make([]byte, x)
		temp,_,_ := r.ReadLine()
		for idx, v := range temp {
			A[i][idx] = v-'0'
		}
	}
	for i := 0; i < y; i++ {
		B[i] = make([]byte, x)
		temp,_,_ := r.ReadLine()
		for idx, v := range temp {
			B[i][idx] = v-'0'
		}
	}
	if x<3 || y<3{
		if check1080(&A,&B){
			fmt.Println(0)
		}else{	
			fmt.Println(-1)
		}
		return
	}
	result :=0
	for i := 0; i < y-2; i++ {
		for j := 0; j < x-2; j++ {
			if A[i][j]!=B[i][j]{
				result+=1
				change1080(&A,j,i)
			}
		}
	}
	if check1080(&A,&B){
		fmt.Println(result)
	}else{	
		fmt.Println(-1)
	}
}

func change1080(graph *[][]byte,x,y int) bool{
	if len(*graph)<y+3{
		return false
	}
	if len((*graph)[0])<x+3{
		return false
	}
	for i := y; i <y+3; i++ {
		for j := x; j < x+3; j++ {
			(*graph)[i][j] = ((*graph)[i][j]+1)%2
		}
	}
	return true
}

func check1080(a,b *[][]byte)bool{
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len((*a)[i]); j++ {
			if (*a)[i][j] != (*b)[i][j]{
				return false
			}
		}
	}
	return true
}