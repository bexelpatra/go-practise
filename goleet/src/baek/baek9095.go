package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek9095() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	N := scanInt9095(r)
	max := 0
	list := make([]int, N)

	for i := 0; i < N; i++ {
		n := scanInt9095(r)
		list[i] = n
		max = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}(max, n)
	}
	dp := make([]int, max+4)

	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i < max+1; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	for _, val := range list {
		fmt.Fprintln(w, dp[val])
	}

}

func scanInt9095(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return num
}
