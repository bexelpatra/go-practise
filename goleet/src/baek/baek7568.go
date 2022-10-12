package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Baek7568() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)

	n := scanInt7568(r)
	people := make([][3]int, n)
	for idx := range people {
		people[idx][0] = scanInt7568(r)
		people[idx][1] = scanInt7568(r)
		people[idx][2] = idx
	}

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] && people[i][1] > people[j][1] {
			return true
		} else if people[i][0] < people[j][0] && people[i][1] < people[j][1] {
			return false
		} else {
			return true
		}
	})
	answer := make([]int, n)

	now := 1
	plus := 0
	answer[people[0][2]] = 1

	for i := 0; i < n-1; i++ {
		j := i + 1
		now += 1
		if people[i][0] > people[j][0] && people[i][1] > people[j][1] {
			answer[people[j][2]] = now
		} else if people[i][0] < people[j][0] && people[i][1] < people[j][1] {
			answer[people[j][2]] = now + plus
			plus = 0
		} else {
			now -= 1
			answer[people[j][2]] = now
			plus += 1
		}
		fmt.Println(answer, plus)
	}
	fmt.Fprintln(w, answer)

}

func scanInt7568(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
