package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2293() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt2293(r)
	target := scanInt2293(r)
	coinVal := make([]int, n)
	dp := make([]int, target+1)
	for i := range coinVal {
		coinVal[i] = scanInt2293(r)
	}
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j >= coinVal[i] {
				dp[j] += dp[j-coinVal[i]]
			}
		}
	}
	fmt.Println(dp[target])
}

func scanInt2293(r *bufio.Scanner) (num int) {
	r.Scan()

	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return
}
