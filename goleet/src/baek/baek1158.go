package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1158() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, k int
	fmt.Fscanln(r, &n, &k)
	circle := make([]int, n+1)
	for i := range circle {
		circle[i] = i + 1
	}
	now := 0
	for len(circle) > 1 {
		now += k
		seq := now % (len(circle) + 1)
		fmt.Println(out(&circle, seq))
	}

}
func out(list *[]int, idx int) int {
	exclude := (*list)[idx]
	var temp []int
	temp = (*list)[:idx]
	temp = append(temp, (*list)[idx+1:]...)
	*list = temp
	return exclude
}
