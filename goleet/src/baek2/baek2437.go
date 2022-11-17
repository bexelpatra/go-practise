package baek2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek2437() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:= scanInt2437(r)
	list := make([]int,n)
	result := make(map[int]struct{},0)
	max :=1
	total :=1
	for i := 0; i < n; i++ {
		list[i] = scanInt2437(r)
		set2437(&result,list[i])
		max*=(i+1)
		total += list[i]
	}
	for i := 1; i < n; i++ {
		visited := make([]bool,n)
		recursive2437(&list,i,0,0,&visited,&result)
	}
	
	// fmt.Println(result)
	for i := 1; i < max; i++ {
		_, ok :=result[i]
		if !ok{
			fmt.Println(i)
			return
		}
	}
	fmt.Println(total)
}

func scanInt2437(r *bufio.Scanner) int {
	r.Scan()
	num:=0
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}

func recursive2437(list *[]int, target, now, value int, visited *[]bool ,result *map[int]struct{}){
	if target == now {
		set2437(result,value)
		return
	}
	for i := 0; i < len(*list); i++ {
		if ((*visited)[i]){
			continue
		}
		(*visited)[i] = true
		value+= (*list)[i]
		recursive2437(list,target,now+1,value,visited,result)
		value-= (*list)[i]
		(*visited)[i] = false
	}	
}

func set2437(set *map[int]struct{}, add int){
	_, ok :=(*set)[add]
	if !ok{
		(*set)[add] = struct{}{}
	}
}


func Baek2437_2(){
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:= scanInt2437(r)
	list := make([]int,n)
	for i := 0; i < n; i++ {
		list[i] = scanInt2437(r)
	}
	sort.Ints(list)
	sum :=0
	for i := 0; i < n; i++ {
		if sum+1< list[i]{
			break
		}
		sum+=list[i]
	}
	fmt.Println(sum+1)
}

func Baek2437_3(){
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:= scanInt2437(r)
	list := make([]int,n)
	
	for i := 0; i < n; i++ {
		list[i] = scanInt2437(r)
	}
	sort.Ints(list)
	dp := make([]int, 10000000)
	sum:=0
	for i := 0; i < n; i++ {
		for j := 1; j < sum+1; j++ {
			if dp[j] == 1{
				dp[j + list[i]] = 1
			}else{
				break
			}
		}
		dp[list[i]] = 1
		sum+=list[i]
	}
	fmt.Println(dp)
}