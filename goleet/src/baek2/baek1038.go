package baek2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Baek1038() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()

	N, _ := strconv.Atoi(r.Text())

	var cnt int64
	check := 0
	list := make([]int64, 0)
	if N > 1022 {
		fmt.Println(-1)
		return
	}
	for check != N+1 {
		if isDecrease(cnt) {
			check++
			list = append(list, int64(cnt))
		}
		cnt += 1
	}
	fmt.Println(list[len(list)-1])
	fmt.Println("cnt! ", cnt)
}

func isDecrease(cnt int64) bool {
	t := strconv.FormatInt(cnt, 10)
	if len(t) < 2 {
		return true
	}
	for i := 0; i < len(t)-1; i++ {
		if t[i] <= t[i+1] {
			return false
		}
	}
	return true
}

func Baek1038_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	list := make([]int64, 0)
	N, _ := strconv.Atoi(r.Text())
	if N < 11 {
		fmt.Println(N)
	} else if N > 1022 {
		fmt.Println(-1)
	} else {
		var i int64
		for i = 0; i < 10; i++ {
			bp(&list, i, 1)
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i] < list[j]
		})
		fmt.Println(list[N])
	}
}
func bp(list *[]int64, num, idx int64) {
	if idx > 10 {
		return
	}
	*list = append(*list, num)
	var i int64
	for i = 0; i < num%10; i++ {
		bp(list, (num*10)+i, idx+1)
	}
}
