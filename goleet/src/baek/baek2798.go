package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek2798() {

	r := bufio.NewReader(os.Stdin)

	var N, target int
	fmt.Fscanln(r, &N, &target)

	temp, _, _ := r.ReadLine()
	tempStr := strings.Split(string(temp), " ")
	cards := make([]int, N+1)
	for i, num := range tempStr {
		cards[i], _ = strconv.Atoi(num)
	}
	max := 0
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				temp := cards[i] + cards[j] + cards[k]
				if temp < target {
					max = func(a, b int) int {
						if a > b {
							return a
						}
						return b
					}(max, temp)
				} else if temp == target {
					fmt.Println(target)
					return
				}
			}
		}
	}
	fmt.Println(max)
}

// func recursive(cardsp *[]int, temp *[]int, now, th int) []int {
// 	if th == 3 {
// 		return *temp
// 	}
// }
