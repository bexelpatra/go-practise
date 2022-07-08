package baek

import (
	"bufio"
	"fmt"
	"os"
)

// st-lab.tistory.com/137 참조
/*
	1 3 5 2 3 4 반례
*/
func Baek11053() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := nextInt11053(r)
	list := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = nextInt11053(r)
	}

	/*
	 */
	//1.  bottom up
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			// fmt.Printf("list[%d] : %d, list[%d] : %d, dp[%d] : %d, dp[%d] : %d \n", i, list[i], j, list[j], i, dp[i], j, dp[j])
			if list[j] < list[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	// 2. top down
	/*
		for i := 0; i < n; i++ {
			recurtiv11053(dp, list, i)
		}
		max := -1
		for i := 0; i < n; i++ {
			max = func(a, b int) int {
				if a < b {
					return b
				}
				return a
			}(max, dp[i])
		}
	*/

	// 3. dfs 시도중...(현재 오답)
	// dfs11053(dp, list, 0, 1, list[0])

	fmt.Println(dp)
	fmt.Println(dp[n-1])
}

func nextInt11053(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return num
}

func recurtiv11053(dp []int, list []int, n int) int {
	if dp[n] == 0 {
		dp[n] = 1

		for i := n - 1; i >= 0; i-- {
			if list[i] < list[n] {
				dp[n] = func(a, b int) int {
					if a < b {
						return b
					}
					return a
				}(dp[n], recurtiv11053(dp, list, i)+1)
			}
		}
	}
	return dp[n]
}

func topDown11053(dp []int, list []int, n int) int {
	if dp[n] == 0 {
		dp[n] = 1
		for i := n - 1; i >= 0; i-- {
			dp[n] = max(dp[n], topDown11053(dp, list, i)+1)
		}
	}
	return dp[n]
}

// 시간 초과로 실패한다.
func dfs11053(dp []int, list []int, idx, count, standard int) {
	if idx > len(list)-1 {
		return
	}
	if idx == len(list)-1 {
		dp[idx] = max11053(dp[idx], count+1)
		return
	}
	if list[idx] > standard {
		dfs11053(dp, list, idx+1, count+1, list[idx])
	} else {
		dfs11053(dp, list, idx+1, count, standard)
		dfs11053(dp, list, idx+1, 1, list[idx])
	}

}

func max11053(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func Baek11053_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := nextInt11053(r)
	list := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = nextInt11053(r)
		dp[i] = 1
	}
	for i := 0; i < n-1; i++ {
		cnt := 1
		standard := list[i]
		for j := i + 1; j < n; j++ {
			if standard < list[j] {
				cnt += 1
				standard = list[j]
				dp[i] = max11053(dp[j], cnt)
			}
		}
	}
	max := 0
	for _, val := range dp {
		max = max11053(val, max)
	}
	fmt.Println(dp)
	fmt.Println(max)
}
