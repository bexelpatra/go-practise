package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Baek11773() {
	r := bufio.NewScanner(os.Stdin)
	n := scanInt11773(r)

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		now := scanInt11773(r)
		if now == 0 && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, now)
		}
	}

	total := 0
	for i, _ := range stack {
		total += stack[i]
	}
	fmt.Println(total)
}

func scanInt11773(r *bufio.Scanner) (number int) {
	r.Scan()
	temp := r.Text()
	number, _ = strconv.Atoi(temp)
	return
}
