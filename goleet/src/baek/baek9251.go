package baek

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Baek9251() {

	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	a := strings.Split(r.Text(), "")
	lenA := len(a)

	r.Scan()
	b := strings.Split(r.Text(), "")
	lenB := len(b)

	dp := make([][]int, lenA+1)
	for i := range dp {
		dp[i] = make([]int, lenB+1)
	}

	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max9251(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}

}
func max9251(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func scanInt9251(r *bufio.Scanner) (num int) {
// 	r.Scan()
// 	for _, b := range r.Bytes() {
// 		num *= 10
// 		num += int(b - '0')
// 	}
// 	return
// }
