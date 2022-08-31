package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek1931() {
	r := bufio.NewReader(os.Stdin)
	var n int
	answer := 0
	fmt.Fscanln(r, &n)

	meetings := make([][2]int64, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &meetings[i][0], &meetings[i][1])
	}
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][1] == meetings[j][1] {
			return meetings[i][0] < meetings[j][0]
		}
		return meetings[i][1] < meetings[j][1]
	})
	now := int64(-1)
	for _, meeting := range meetings {
		if meeting[0] >= int64(now) {
			now = meeting[1]
			answer += 1
		}
	}
	fmt.Println(answer)
}

func Baek1931_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt1931(r)

	answer := 0
	meatings := make([][2]int, n)

	// for index, _ := range meatings {
	// 	// 1. call by value 인 경우 indexing 해서 객체를 slice를 가져와야 한다.
	// 	meating := &meatings[index]
	// 	meating[0], meating[1] = scanInt1931(r), scanInt1931(r)

	// 	// 2. call by refence 인 경우 값을 그대로 넣어주면 meatings에 값이 들어가 있어야 한다.
	// 	// value[0],value[1] = scanInt1931(r), scanInt1931(r)
	// }
	sort.Slice(meatings, func(i, j int) bool {
		if meatings[i][1] == meatings[j][1] {
			return meatings[i][0] < meatings[j][0]
		}
		return meatings[i][1] < meatings[j][1]
	})

	fmt.Println(meatings)
	now := 0
	for _, meating := range meatings {
		if meating[0] > now {
			now = meating[1]
			answer += 1
		}
	}
	fmt.Println(answer)
}

func scanInt1931(r *bufio.Scanner) int {
	result := 0
	r.Scan()
	for _, val := range r.Bytes() {
		result *= 10
		result += int(val - '0')
	}
	return result
}
