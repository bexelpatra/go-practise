package baek

import (
	"bufio"
	"fmt"
	"os"
)
var(
	N int
	K int
)
func Baek14502() {
	r:= bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	N = scanInt14502(r)
	K = scanInt14502(r)
	graphs,virus,walls := scanLine14502(r,N,K)

	max :=0
	for i := 0; i < len(walls)-2; i++ {
		wall1 :=walls[i]
		for j := i+1; j < len(walls)-1; j++ {
			wall2:=walls[j]
			for k := j+1; k < len(walls); k++ {
				wall3:=walls[k]

				checked := make([][]bool,N)
				for v := range checked {
					checked[v] = make([]bool, K)
				}
				changGraph(&graphs,wall1,wall2,wall3,1)
				max = func(a,b int)int{
					if(a>b){
						return a
					}
					return b
				}(spreadVirus(graphs,checked,virus),max)
				changGraph(&graphs,wall1,wall2,wall3,0)
			}
		}
	}
	fmt.Println(max)
	// for _, v := range graphs {
	// 	fmt.Println(v)
	// }
}
func changGraph(graph *[][]int, w1,w2,w3 []int,to int){
	(*graph)[w1[0]][w1[1]] = to
	(*graph)[w2[0]][w2[1]] = to
	(*graph)[w3[0]][w3[1]] = to
}

func scanLine14502(r *bufio.Scanner,n,k int)([][]int,[][2]int,[][]int){
	result := make([][]int,n)
	virus := make([][2]int,0)
	walls :=make([][]int,0)
	for i := range result {
		result[i] = make([]int, k)
		for j := range result[i] {
			result[i][j] = scanInt14502(r)
			if(result[i][j] ==2 ){
				virus = append(virus, [2]int{i,j})
			}else if result[i][j] ==0{
				walls = append(walls, []int{i,j})
			}
		}
	}
	return result,virus,walls
}

func scanInt14502(r * bufio.Scanner) int{
	r.Scan()
	num:=0
	for _, v := range r.Bytes() {
		num*=10
		num +=(int)(v-'0')
	}
	return num
}

func spreadVirus(graph [][]int, check [][]bool, virus [][2]int)int{
	num :=0
	tempGraph :=copy14502(graph)
	for _, v := range virus {
		recursive14502(tempGraph,check,v[0],v[1])
	}
	// for _, v := range tempGraph {
	// 	fmt.Println(v)
	// }
	// for _, v := range graph {
	// 	fmt.Println(v)
	// }
	num = countZero14502(tempGraph)
	return num
}

func recursive14502(graph [][]int,check [][]bool,n,k int){
	if(n>=N || k >=K){
		return
	}
	if(n<0 || k<0){
		return
	}
	if(check[n][k]){
		return
	}
	check[n][k]=true
	if (graph[n][k]==0){
		graph[n][k]=2
	}
	switch (graph)[n][k]{
		case 0: {
		}
		case 1:{

		}
		case 2:{
			recursive14502(graph,check,n-1,k)
			recursive14502(graph,check,n,k-1)
			recursive14502(graph,check,n+1,k)
			recursive14502(graph,check,n,k+1)
		}
		default :{
			fmt.Println("이런일은 안생김 ㅋ")
		}
	}
}

func copy14502(graph [][]int) (copied [][]int){
	copied =make([][]int,len(graph))
	for i, v := range graph {
		copied[i] = make([]int, len(v))
		for j, v2 := range v {
			copied[i][j]=v2
		}
	}
	return
}

func countZero14502(graph [][]int)int{
	num:=0
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if(graph[i][j]==0)	{
				num+=1
			}
		}
	}
	return num
}