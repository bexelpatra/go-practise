package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek1541() {
	r := bufio.NewReader(os.Stdin)

	b, _, _ := r.ReadLine()
	plus := 0
	p := ""
	minus := 0
	m := ""
	flag := false

	for _, char := range strings.Split(string(b), "") {
		if char == "+" {
			numP, _ := strconv.Atoi(p)
			numM, _ := strconv.Atoi(m)
			p = ""
			m = ""
			plus += numP
			minus += numM
		} else if char == "-" {
			flag = true
			numP, _ := strconv.Atoi(p)
			numM, _ := strconv.Atoi(m)
			p = ""
			m = ""
			plus += numP
			minus += numM
		} else {
			if !flag {
				p += char
			} else {
				m += char
			}
		}
	}

	numP, _ := strconv.Atoi(p)
	numM, _ := strconv.Atoi(m)
	p = ""
	m = ""
	plus += numP
	minus += numM

	fmt.Println(plus - minus)
}
