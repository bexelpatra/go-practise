package baek

import (
	"bufio"
	"fmt"
	"os"
)

// dp[n] = dp[n - arr[i]] +1
func Baek1912() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt1912(r)
	numbers := make([]int, n)
	dp := make([]int, n)
	for i := range numbers {
		numbers[i] = scanInt1912(r)
	}

	dp[0] = numbers[0]

	for i := 1; i < len(numbers); i++ {
		if dp[i-1] < 0 {
			dp[i] = func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}(dp[i-1], numbers[i])
		} else {
			dp[i] = dp[i-1] + numbers[i]
		}
	}
	max := -1001
	for _, v := range dp {
		max = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}(max, v)
	}
	fmt.Println(max)

}

func scanInt1912(r *bufio.Scanner) (num int) {
	r.Scan()
	minus := false
	for _, b := range r.Bytes() {
		num *= 10
		if b-'0' > '9' {
			minus = true
		} else {
			num += int(b - '0')
		}
	}
	if minus {
		num = -num
	}
	return
}
