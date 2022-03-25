package baek

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)

	var N, target int
	fmt.Fscanln(r, &N, &target)

	dif := N - target
	dif = dif * dif
	if target < N {
		N, target = target, N
	}
	if dif == 0 {
		fmt.Println("0")
		return
	}
	lines := make([]int, N*N+1)
	for i, _ := range lines {
		lines[i] = 10000
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

	fmt.Println(lines[target])
}
