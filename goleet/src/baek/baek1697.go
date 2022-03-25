package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1697() {
	// dp와 같은 방식으로 풀이하려고 했지만 실패
	// 시간초과
	r := bufio.NewReader(os.Stdin)

	var N, target int
	fmt.Fscanln(r, &N, &target)
	dif := N - target
	dif = dif * dif
	if target < N {
		N, target = target, N
	}
	if dif == 0 {
		fmt.Println(0)
		return
	}
	lines := make([]int, dif+N)
	for i, _ := range lines {
		lines[i] = 1000
	}
	if N > 0 {
		lines[N-1] = 1
	}
	lines[N] = 0
	lines[N+1] = 1
	lines[N*2] = 1
	for i := N + 2; i < len(lines)-1; i++ {
		if i/2 == 1 {
			lines[i] = min(lines[i-1], lines[i+1]) + 1
		}
		if i%2 == 0 {
			lines[i] = min(lines[i-1], min(lines[i+1], lines[i/2])) + 1
		} else {
			lines[i] = min(lines[i-1], min(lines[i+1], lines[i/2]+1)) + 1

		}
	}
	fmt.Println(lines)
	fmt.Println(lines[target])

}
func findTarget(lines []int, now int) int {
	if now < 0 {
		return lines[now+1] + 1
	}
	if len(lines) <= now {
		return 0
	}
	if lines[now] == 1000 {
		a := findTarget(lines, now-1)
		b := findTarget(lines, now+1)
		c := findTarget(lines, now*2)

		lines[now] = min(a, min(b, c))

	}

	return lines[now]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Baek11697_2() {
	r := bufio.NewReader(os.Stdin)

	var N, target int
	fmt.Fscanln(r, &N, &target)

	if N > target {
		N, target = target, N
	}
	lines := make([]int, 100000+1)
	if N > 0 {
		lines[N-1] = 1
	}
	lines[N] = 0
	lines[N+1] = 1
	lines[N*2] = 1
	for {

	}
}
