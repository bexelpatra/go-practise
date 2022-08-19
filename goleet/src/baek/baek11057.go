package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek11057() {
	r := bufio.NewScanner(os.Stdin)
	n := scanInt11057(r)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, 10)
		dp[i][0] = 1
	}
	for i := 0; i < 10; i++ {
		dp[1][i] = 1
	}

	result := 1
	if n == 1 {
		result = 10
	}
	for i := 2; i < n+1; i++ {
		for j := 1; j < 10; j++ {
			dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % 10007
			if i == n {
				result += dp[i][j]
			}
		}
	}

	fmt.Println(result % 10007)

}

func scanInt11057(r *bufio.Scanner) (num int) {
	r.Scan()
	for _, b := range r.Bytes() {
		num *= 10
		num += int(b - '0')
	}
	return
}
