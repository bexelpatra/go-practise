package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Baek11726() {

	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	dp := make([]int, n+3)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10007
	}
	fmt.Println(dp[n])
}

func Baek11726_2() {

	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	dp := make([]int, n+3)
	dp[1] = 1
	dp[2] = 2

	topDown11726(&dp, n)

	fmt.Println(dp[n])
}
func topDown11726(dp *[]int, idx int) int {
	if (*dp)[idx] == 0 {
		(*dp)[idx] = (topDown11726(dp, idx-1) + topDown11726(dp, idx-2)) % 10007
	}
	return (*dp)[idx]

}
