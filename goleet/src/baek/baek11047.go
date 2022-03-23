package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek11047() {
	r := bufio.NewReader(os.Stdin)
	var n int
	var target int
	var answer int
	fmt.Fscanln(r, &n, &target)
	coins := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		fmt.Fscanln(r, &coins[i])
	}
	for _, coin := range coins {

		if target < coin {
			continue
		}
		temp := target % coin
		answer += (target / coin)
		target = temp
		if temp == 0 {
			break
		}
	}
	fmt.Println(answer)
}
