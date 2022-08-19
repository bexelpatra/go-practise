package baek

import (
	"bufio"
	"fmt"
	"os"
)

// 1로 만들기
func Baek1463() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = Min1464(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = Min1464(dp[i], dp[i/3]+1)
		}
	}
	fmt.Println(dp[n])
}

func Baek1463_2() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = Min1464(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = Min1464(dp[i], dp[i/3]+1)
		}
	}
	fmt.Println(dp[n])
}

func Baek1463_3() {
	r := bufio.NewScanner(os.Stdin)
	n := scanInt1463(r)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = Min1464(dp[i/2]+1, dp[i])
		}
		if i%3 == 0 {
			dp[i] = Min1464(dp[i/3]+1, dp[i])
		}
	}
	fmt.Println(dp[n])
}

func Baek1463_4() {
	r := bufio.NewScanner(os.Stdin)
	n := scanInt1463(r)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = 100000
	}
	TopDown1464(dp, n)
	fmt.Println(dp)
}
func Min1464(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TopDown1464(dp []int, idx int) int {
	if dp[idx] < 100000 {
		return dp[idx]
	}
	if idx == 0 || idx == 1 {
		dp[idx] = 0
		return 0
	}
	if idx == 2 || idx == 3 {
		dp[idx] = 1
		return dp[idx]
	}
	dp[idx] = Min1464(TopDown1464(dp, idx-1)+1, dp[idx])
	if idx%2 == 0 {
		dp[idx] = Min1464(TopDown1464(dp, idx/2)+1, dp[idx])
	}
	if idx%3 == 0 {
		dp[idx] = Min1464(TopDown1464(dp, idx/3)+1, dp[idx])
	}

	return dp[idx]
}

func scanInt1463(r *bufio.Scanner) int {
	num := 0
	r.Scan()
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return num
}
