package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Baek1026() {

	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(r, &n)
	rows := make([][]string, 2)
	for i := 0; i < 2; i++ {
		temp, _, _ := r.ReadLine()
		rows[i] = strings.Split(string(temp), " ")
	}
	a := convert1026(rows[0])
	b := convert1026(rows[1])

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	total := 0

	for i := 0; i < n; i++ {
		total += a[i] * b[i]

	}

	fmt.Println(total)

}
func convert1026(str []string) (result []int) {
	result = make([]int, len(str))
	for i, num := range str {
		temp, _ := strconv.Atoi(num)
		result[i] = temp
	}
	return
}
