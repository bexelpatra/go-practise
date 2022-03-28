package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek5585() {
	r := bufio.NewReader(os.Stdin)

	total := 0
	var paid int
	fmt.Fscanln(r, &paid)

	amt := 1000 - paid

	changes := []int{500, 100, 50, 10, 5, 1}

	for _, change := range changes {
		if amt == 0 {
			break
		}
		temp := amt / change
		rest := amt % change
		total += temp
		if temp > 0 {
			amt = rest
		}
	}

	fmt.Println(total)

}
