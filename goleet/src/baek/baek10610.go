package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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
	sort.Slice(numbers,func(i, j int) bool {return i>j})
	var s s
	s = strings.Split("1238740","")
	
	
	sort.Slice(s, func(i, j int) bool {
		return s[i]>s[j]
	})
	
	fmt.Println(s)
	sort.Sort(s)
	fmt.Println(s)
	
}
type s []string
func (s s) Len() int           { return len(s) }
func (s s) Less(i, j int) bool { return s[i] > s[j] }
func (s s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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