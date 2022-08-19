package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Baek10989() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt10989(r)
	list := make([]int, 0)
	for i := 0; i < n; i++ {
		list = append(list, scanInt10989(r))
	}
	sort.Ints(list)
	for i, _ := range list {
		fmt.Fprintln(w, list[i])
	}

}

func scanInt10989(r *bufio.Scanner) int {
	r.Scan()
	temp, _ := strconv.Atoi(r.Text())
	return temp
}
func Baek10989_2() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := scanInt10989(r)
	list := make([]int, 10001)
	for i := 0; i < n; i++ {
		list[scanInt10989(r)] += 1
	}

	for i, _ := range list {
		if list[i] != 0 {
			for j := 0; j < list[i]; j++ {
				fmt.Fprintln(w, i)
			}
		}
	}

}
