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
