package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// baek.Baek2606()
	// baek.Baek2606_2()
	// baek.Baek1012()
	// baek.Baek1717()
	r := bufio.NewReader(os.Stdin)
	var str string
	var len int
	fmt.Fscanln(r, &len)
	fmt.Fscanln(r, &str)

	// len := len(str)
	total := 0
	for _, val := range str {
		total += int(val) - '0'
	}
	fmt.Println(total)
}
