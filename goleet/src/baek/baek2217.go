package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek2217() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(r,&n)
	lines := make([]int,n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r,&lines[i])
	}
	sort.Slice(lines,func (a,b int) bool  {
		return lines[a]>lines[b]
	})
	max := lines[0]
	for i := 1; i < n; i++ {
		max = func(a,b int) int{
			if a <b{
				return b
			}
			return a
		}(max,lines[i] * (i+1))
	}
	fmt.Println(max)
}