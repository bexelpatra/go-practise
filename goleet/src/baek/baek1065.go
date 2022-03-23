package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Baek1065() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	total := 0

	if n < 100 {
		fmt.Println(n)
		return
	}

	total += 99
	for i := 100; i <= n; i++ {
		flag := true
		temp := strconv.Itoa(i)
		dif := (int(temp[0]) - int(temp[1]))
	label:
		for j := 1; j < len(temp)-1; j++ {
			if dif != (int(temp[j]) - int(temp[j+1])) {
				flag = false
				break label
			}
		}
		if flag {
			total += 1
		}

	}
	fmt.Println(total)

}
