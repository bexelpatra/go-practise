package baek

import (
	"bufio"
	"fmt"
	"os"
)

// dp[n] = dp[n - arr[i]] +1
func Baek2294() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt2294(r)
	k := scanInt2294(r)
	coin := make(map[int]struct{})
	dp := make([]int, k+1)

	for i := 0; i < n; i++ {
		coin[scanInt2294(r)] = struct{}{}
	}

	for i := range dp {
		dp[i] = 100_000_000
	}
	dp[0] = 0
	for i := 1; i < k+1; i++ {
		for key := range coin {
			if i == key {
				dp[i] = key
			}
			if i >= key {
				dp[i] = func(a, b int) int {
					if a > b {
						return b
					}
					return a
				}(dp[i], dp[i-key]+1)
			}
		}
	}
	for i, v := range dp {
		fmt.Println(i, " : ", v)
	}
	if dp[k] == 100000000 {
		dp[k] = -1
	}
	fmt.Println(dp[k])

}

func scanInt2294(r *bufio.Scanner) (num int) {
	r.Scan()

	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return
}
