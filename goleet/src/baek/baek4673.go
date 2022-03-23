package baek

import (
	"fmt"
	"strconv"
)

var (
	visit   []bool
	numbers []int
)

func Baek4673() {
	visit = make([]bool, 10002)
	numbers = make([]int, 10502)

	for i := 1; i < 10001; i++ {
		d(i)
	}

	for i := 1; i < 10000; i++ {
		if numbers[i] == 0 {
			fmt.Println(i)
		}
	}
}

func d(num int) {
	if num > 10001 {
		return
	}
	if visit[num] {
		return
	}
	visit[num] = true

	str := strconv.Itoa(num)

	for _, n := range str {
		num += (int(n) - '0')
	}
	numbers[num] -= 1
	d(num)
}
