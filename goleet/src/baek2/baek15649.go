package baek2

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Baek15649() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n:=scanInt15649(r)
	num:=scanInt15649(r)
	list := make([]int,n)
	total :=1
	for i := 0; i < n; i++ {
		list[i] = i+1
		total *=(n-i)
	}
	sb := bytes.Buffer{}
	sb2:=bytes.Buffer{}
	result15649=make([]string, 0)
	recursive15649_2(&list,num,0,&sb)
	
	// fmt.Println(sb.String())
	sort.Strings(result15649)
	for _, v := range result15649 {
		sb2.WriteString(v)
	}
	// fmt.Println(sb2.String())
	visited :=make([]bool,n+1)
	result :=make([]string,0)
	perm := make([]int,num)
	useVisited(&list,&perm,&visited,num,0,&result)
	fmt.Println(result)
}

func scanInt15649(r *bufio.Scanner)int {
	num:=0
	r.Scan()
	for _, v := range r.Bytes() {
		num*=10
		num+=(int)(v-'0')
	}
	return num
}
var(
	result15649 []string
)
func sliceString15649(s *[]int)string{
	sb := bytes.Buffer{}
	for _, v := range (*s) {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(" ")
	}
	// sb.WriteString("\n")
	return sb.String()
}

func recursive15649_2(list *[]int ,target,now int, sb *bytes.Buffer){
	if target == now{
		temp :=(*list)[:target]
		sb.WriteString(sliceString15649(&temp))
		sb.WriteString("\n")
		result15649 = append(result15649, sliceString15649(&temp))
		return
	}
	for i := now; i < len(*list); i++ {
		swap15649(list,now,i)
		recursive15649_2(list,target,now+1,sb)
		swap15649(list,now,i)
	}
}

func swap15649(list *[]int,a,b int){
	(*list)[a],(*list)[b] = (*list)[b],(*list)[a]
}
/*
private static void makePermutation(int r, int[] temp, int current, boolean[] visited) {
		if (r == current) {
			System.out.println(Arrays.toString(temp));
		} else {
			for (int i = 0; i < arr.length; i++) {
				if (!visited[i]) {
					visited[i] = true;
					temp[current] = arr[i];
					makePermutation(r, temp, current +1, visited);
					visited[i] = false;
				}
			}
		}
	}
*/

func useVisited(list,temp *[]int,visited *[]bool,target, now int, result *[]string){
	if target == now{
		sb := bytes.Buffer{}
		for _, v := range *temp {
			sb.WriteString(strconv.Itoa(v))
		}
		(*result) = append((*result), sb.String())
		return
	}else{

	}
	for i := 0; i < len(*list); i++ {
		if (*visited)[i]{
			continue
		}
		(*visited)[i] = true
		(*temp)[now] = (*list)[i]
		useVisited(list,temp,visited,target,now+1,result)
		(*visited)[i] = false
	}
}