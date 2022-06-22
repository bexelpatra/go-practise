package baek

import (
	"bufio"
	"fmt"
	"os"
)

// 가방챙기기
// 점화식 세우기가 어렵다...
func Baek12865() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	N := nextInt12865(r)
	K := nextInt12865(r)

	w := make([]int, N+1)
	v := make([]int, N+1)
	dp := make([][]int, N+1)
	for i := 1; i < N+1; i++ {
		w[i] = nextInt12865(r)
		v[i] = nextInt12865(r)
	}
	dp[0] = make([]int, K+1)
	for i := 1; i <= N; i++ {
		dp[i] = make([]int, K+1)

		for j := 1; j <= K; j++ {
			if w[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = func(a, b int) int {
					if a > b {
						return a
					}
					return b
				}(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}

	}
	fmt.Println(dp[N][K])
}

func Baek12865_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	N := nextInt12865(r)
	K := nextInt12865(r)

	maxValue := 0
	dp := make([]int, K+1)

	for ; N > 0; N-- {
		w, v := nextInt12865(r), nextInt12865(r)
		for i := K; i >= w; i-- {
			if dp[i] == 0 || dp[i] < dp[i-w]+v {
				dp[i] = dp[i-w] + v
				if dp[i] > maxValue {
					maxValue = dp[i]
				}
			}
		}
	}
	fmt.Println(dp[K])
	fmt.Println(dp)
}

func nextInt12865(r *bufio.Scanner) int {
	num := 0
	r.Scan()
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return num
}
