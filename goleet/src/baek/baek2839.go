package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2839() {

	r := bufio.NewReader(os.Stdin)
	var target int
	fmt.Fscanln(r, &target)
	if target == 3 || target == 5 {
		fmt.Println(1)
		return
	}
	now := []int{3, 5}
	count := 1
	length := len(now)

	for {
		for i, _ := range now {
			now[i] += 3

			if now[i] == target {
				fmt.Println(count + 1)
				return
			}
		}
		now = append(now, now[length-1]+2)

		if now[length] == target {
			fmt.Println(count + 1)
			return
		}
		if now[0] > target {
			fmt.Println(-1)
			return
		}
		length += 1
		count += 1
	}

}
