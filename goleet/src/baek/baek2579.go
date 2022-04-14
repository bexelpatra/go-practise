package baek

import (
	"bufio"
	"fmt"
	"os"
)

/*
6
10
20
15
25
10
20
*/
func Baek2579() {
	r := bufio.NewScanner(os.Stdin)
	n := scanInt2579(r)
	list := make([]int, n+1)
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		list[i] = scanInt2579(r)
	}

	if n == 1 {
		fmt.Println(list[1])
		return
	}
	if n == 2 {
		fmt.Println(list[1] + list[2])
		return
	}
	dp[1] = list[1]
	dp[2] = list[1] + list[2]
	if n < 3 {
		sum := 0
		for i := 1; i < n+1; i++ {
			sum += list[n]
		}
		fmt.Println(sum)
		return
	}
	for i := 3; i < n+1; i++ {
		dp[i] = func(a, b int) int {
			if a < b {
				return b
			}
			return a
		}(dp[i-2]+list[i], dp[i-3]+list[i-1]+list[i])

	}
	fmt.Println(dp[n])
}

func scanInt2579(r *bufio.Scanner) (num int) {
	r.Scan()
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return
}
