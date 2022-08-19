package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1149() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt1149(r)
	list := make([]map[string]int, n+1)
	dp := make([]map[string]int, n+1)

	for i := 0; i < n; i++ {
		list[i] = scanLine1149(r)
	}

	dp[1] = make(map[string]int)
	dp[1]["r"] = list[1]["r"] + min1149(list[0]["g"], list[0]["b"])
	dp[1]["g"] = list[1]["g"] + min1149(list[0]["r"], list[0]["b"])
	dp[1]["b"] = list[1]["b"] + min1149(list[0]["r"], list[0]["g"])

	for i := 2; i < n+1; i++ {
		dp[i] = make(map[string]int)
		dp[i]["r"] = list[i]["r"] + min1149(dp[i-1]["g"], dp[i-1]["b"])
		dp[i]["g"] = list[i]["g"] + min1149(dp[i-1]["r"], dp[i-1]["b"])
		dp[i]["b"] = list[i]["b"] + min1149(dp[i-1]["r"], dp[i-1]["g"])
	}

	fmt.Println(min1149(dp[n]["r"], min1149(dp[n]["g"], dp[n]["b"])))

}

func scanInt1149(r *bufio.Scanner) int {
	num := 0
	r.Scan()
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return num
}

func scanLine1149(r *bufio.Scanner) map[string]int {
	s := make(map[string]int, 3)
	rgb := []string{"r", "g", "b"}
	for _, val := range rgb {
		s[val] = scanInt1149(r)
	}
	return s
}

func min1149(a, b int) int {
	if a > b {
		return b
	}
	return a
}
