package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Baek2750() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	numbers := make([]int, 2002)
	defer w.Flush()
	n := scanInt2750(r)

	for i := 0; i < n; i++ {
		numbers[scanInt2750(r)+1_000] = 1
	}

	for i := range numbers {
		if numbers[i] == 1 {
			fmt.Fprintln(w, i-1_000)
		}
	}

}

func Baek2750_2() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	numbers := make([]int, 0)
	defer w.Flush()
	n := scanInt2750(r)

	for i := 0; i < n; i++ {
		numbers = append(numbers, scanInt2750(r))
	}
	sort.Ints(numbers)
	for i := range numbers {
		fmt.Fprintln(w, numbers[i])
	}

}

func scanInt2750(r *bufio.Scanner) int {
	r.Scan()
	temp, _ := strconv.Atoi(r.Text())
	return temp
}
