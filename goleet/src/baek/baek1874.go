package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Baek1874() {
	r := bufio.NewScanner(os.Stdin)
	N := scanInt1874(r)
	s := Stack1874{}
	sb := bytes.Buffer{}
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt1874(r)
	}
	for i := 0; i < N-1; i++ {
		arr[i]
	}

	fmt.Println()
	fmt.Println(sb.String())
}

// sb.WriteString(strconv.Itoa(s.pop()))
// sb.WriteString("\n")
type Stack1874 []int

func (s *Stack1874) push(add int) {
	(*s) = append((*s), add)
}
func (s *Stack1874) pop() int {
	len := len((*s))
	if len < 1 {
		return 0
	}
	result := (*s)[len-1]
	(*s) = (*s)[:len-1]
	return result
}
func (s *Stack1874) peek() int {
	len := len((*s))

	return (*s)[len-1]
}
func scanInt1874(r *bufio.Scanner) int {
	num := 0
	r.Scan()
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
