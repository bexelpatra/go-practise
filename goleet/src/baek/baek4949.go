package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Baek4949() {
	r := bufio.NewScanner(os.Stdin)
	sb := bytes.Buffer{}

	for {
		r.Scan()
		s := make([]int, 0)
		bytes := r.Bytes()
		if bytes[0] == '.' {
			break
		}
		flag := false
		for idx := range bytes {
			if bytes[idx] == '[' {
				s = append(s, 1)
			} else if bytes[idx] == '{' {
				s = append(s, 2)
			} else if bytes[idx] == '(' {
				s = append(s, 3)
			} else if bytes[idx] == ']' {
				if pop4949(&s)-1 != 0 {
					flag = true
					break
				}
			} else if bytes[idx] == '}' {
				if pop4949(&s)-2 != 0 {
					flag = true
					break
				}
			} else if bytes[idx] == ')' {
				if pop4949(&s)-3 != 0 {
					flag = true
					break
				}
			}
		}
		if len(s) > 0 || flag {
			sb.WriteString("no\n")
		} else {
			sb.WriteString("yes\n")
		}

	}
	fmt.Println(sb.String())
}

func pop4949(s *[]int) int {
	if len(*s) < 1 {
		return -1
	}
	temp := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return temp
}
