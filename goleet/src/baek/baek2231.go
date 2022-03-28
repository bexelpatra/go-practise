package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Baek2231() {
	r := bufio.NewReader(os.Stdin)
	var target int
	fmt.Fscanln(r, &target)

	for i := 1; i < target; i++ {
		if i+sum2231(i) == target {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("0")
}
func sum2231(n int) (x int) {

	for _, num := range strconv.Itoa(n) {
		temp := int(num - '0')
		x += temp

	}
	return
}
