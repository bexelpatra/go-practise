package baek2

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type item17298 struct {
	V, i int
}

func Baek17298() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt17298(r)
	list := make([]item17298, n)
	result := make([]int, n)
	stack := make([]item17298, 0)

	list[0] = item17298{scanInt17298(r), 0}
	sb := bytes.Buffer{}
	stack = append(stack, list[0])

	for i := 1; i < n; i++ {
		list[i] = item17298{scanInt17298(r), i}

		p := Peek17298_2(&stack)
		if 0 < p.V && p.V < list[i].V {
			for Peek17298_2(&stack).V < list[i].V && Peek17298_2(&stack).V > 0 {
				result[Pop17298_2(&stack).i] = list[i].V
			}
		}
		if list[i].V == 9 {
			result[i] = -1
		} else {
			stack = append(stack, list[i])
		}
	}

	result[n-1] = -1
	Pop17298_2(&stack)

	for Peek17298_2(&stack).V > 0 {
		result[Pop17298_2(&stack).i] = -1
	}

	for i := 0; i < n; i++ {
		sb.WriteString(strconv.Itoa(result[i]))
		sb.WriteString(" ")
	}
	fmt.Println(sb.String())
}
func Peek17298(stack *[]int) int {
	l := len(*stack)
	if l == 0 {
		return -1
	}
	return (*stack)[l-1]
}
func Pop17298(stack *[]int) int {
	if len(*stack) == 0 {
		return -1
	}
	temp := (*stack)[len(*stack)-1]
	(*stack) = (*stack)[:len(*stack)-1]
	return temp
}

func Peek17298_2(stack *[]item17298) item17298 {
	l := len(*stack)
	if l == 0 {
		return item17298{}
	}
	return (*stack)[l-1]
}
func Pop17298_2(stack *[]item17298) item17298 {
	if len(*stack) == 0 {
		return item17298{}
	}
	temp := (*stack)[len(*stack)-1]
	(*stack) = (*stack)[:len(*stack)-1]
	return temp
}
func scanInt17298(r *bufio.Scanner) int {
	num := 0
	r.Scan()
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
