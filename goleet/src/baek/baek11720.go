package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek11720() {

	r := bufio.NewReader(os.Stdin)
	var str string
	var len int
	fmt.Fscanln(r, &len)
	fmt.Fscanln(r, &str)

	total := 0
	for _, val := range str {
		total += int(val) - '0'
	}
	fmt.Println(total)
}
