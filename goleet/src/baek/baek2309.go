package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// 난쟁이
func Baek2309() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := 9

	list := make([]int, n)
	total := 0
	for i := range list {
		list[i] = scanInt2309(r)
		total += list[i]
	}

	var a, b int

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if total-list[i]-list[j] == 100 {
				a = i
				b = j
			}
		}
	}

	result := make([]int, 0)
	for i := range list {
		if a == i || b == i {
			continue
		}
		result = append(result, list[i])
	}
	sort.Ints(result)
	for i := range result {
		fmt.Fprintln(w, result[i])
	}
}

func scanInt2309(r *bufio.Scanner) int {
	r.Scan()
	temp, _ := strconv.Atoi(r.Text())
	return temp
}
