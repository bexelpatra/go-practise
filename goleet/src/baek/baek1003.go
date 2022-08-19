package baek

import (
	"bufio"
	"fmt"
	"os"
)

// 피보나치 0, 1 개수세기
func Baek1003() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := nextInt1003(r)
	list := make([]int, n)
	max := 0

	for i := 0; i < n; i++ {
		list[i] = nextInt1003(r)
		max = func(a, b int) int {
			if a < b {
				return b
			}
			return a
		}(max, list[i])
	}

	dp := make([][]int, max+1)
	dp[0] = []int{1, 0}
	dp[1] = []int{0, 1}

	for i := 2; i < max+1; i++ {
		dp[i] = []int{dp[i-1][0] + dp[i-2][0], dp[i-1][1] + dp[i-2][1]}
	}

	for _, num := range list {
		fmt.Fprintln(w, dp[num][0], dp[num][1])
	}

}

func Baek1003_2() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := nextInt1003(r)
	list := make([]int, n)
	max := 0

	for i := 0; i < n; i++ {
		list[i] = nextInt1003(r)
		max = func(a, b int) int {
			if a < b {
				return b
			}
			return a
		}(max, list[i])
	}

	dp := make([][]int, max+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	fibo(&dp, max)
	for _, num := range list {
		fmt.Fprintln(w, dp[num][0], dp[num][1])
	}

}

func nextInt1003(r *bufio.Scanner) (num int) {
	r.Scan()
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return
}

func fibo(dp *[][]int, a int) (int, int) {
	if (*dp)[a][0] != 0 && (*dp)[a][1] != 0 {
		return (*dp)[a][0], (*dp)[a][1]
	}
	if a == 0 {
		(*dp)[a][0] = 1
		(*dp)[a][1] = 0
	} else if a == 1 {
		(*dp)[a][0] = 0
		(*dp)[a][1] = 1
	} else {
		a0, a1 := fibo(dp, a-1)
		b0, b1 := fibo(dp, a-2)
		(*dp)[a][0] = a0 + b0
		(*dp)[a][1] = a1 + b1
	}

	return (*dp)[a][0], (*dp)[a][1]
}
