package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func Baek4963() {
	r:=bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	result := bytes.Buffer{}
	for {
		h:= scanInt4963(r)
		w:= scanInt4963(r)
		if w==0 && h ==0{
			break
		}
		graphs:=make([][]int,w)
		checked := make([][]bool,w)
		count:=0
		for idx := range graphs {
			graphs[idx] = make([]int, h)
			checked[idx] = make([]bool, h)
		}

		for i := range graphs {
			for j := range graphs[i] {
				graphs[i][j] = scanInt4963(r)
			}
		}
		for i := range graphs {
			for j := range graphs[i] {
				if recursive4963(&graphs,&checked,w,h,i,j){
					count+=1
				}
			}
		}
		result.WriteString(strconv.Itoa(count))
		result.WriteString("\n")

	}
	fmt.Println(result.String())

}

func recursive4963(graph *[][]int, checked *[][]bool, w,h,x,y int) (flag bool){
	if w -1<  x || h-1 < y{
		return
	}
	if x<0 || y<0{
		return
	}
	if (*checked)[x][y]{
		return
	}
	if((*graph)[x][y]==0){
		return
	}
	
	(*checked)[x][y] = true
	recursive4963(graph,checked,w,h,x-1,y)
	recursive4963(graph,checked,w,h,x,y-1)
	recursive4963(graph,checked,w,h,x+1,y)
	recursive4963(graph,checked,w,h,x,y+1)
	
	recursive4963(graph,checked,w,h,x-1,y-1)
	recursive4963(graph,checked,w,h,x-1,y+1)
	recursive4963(graph,checked,w,h,x+1,y-1)
	recursive4963(graph,checked,w,h,x+1,y+1)
	
	flag=true
	return
}

func scanInt4963(r *bufio.Scanner) int{
	r.Scan()
	num:=0
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}