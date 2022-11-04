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
	numbers := Stack1874{}
	sb := bytes.Buffer{}
	arr := make([]int, N+1)
	arr[0] = 0
	for i := 1; i < N+1; i++ {
		arr[i] = scanInt1874(r)
		numbers.push(N - i + 1)
	}

	for i := 0; i < N; i++ {
		if arr[i] < arr[i+1] {
			for numbers.peek() <= arr[i+1] {
				s.push(numbers.pop())
				sb.WriteString("+")
				sb.WriteString("\n")
			}
		} else {
			if s.peek() != arr[i+1] {
				fmt.Println("NO")
				return
			}
		}
		s.pop()
		sb.WriteString("-")
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

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
	if len < 1 { // 없으면 25번째 줄에서 에러 발생한다.
		return 100_001
	}
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

// https://www.acmicpc.net/source/19232035
// 깔끔하고 빠르다.
func Baek1874_2() {
	r := bufio.NewScanner(os.Stdin)
	n := scanInt1874(r)
	result := make([]byte, 0, n*4)
	s := make([]int, 0, n)
	for i, c := 0, 1; i < n; i++ {
		x := scanInt1874(r)
		for x >= c {
			s = append(s, c)
			result = append(result, "+\n"...)
			c += 1
		}
		if s[len(s)-1] != x {
			os.Stdout.WriteString("NO")
			return
		}
		result = append(result, "-\n"...)
		s = s[:len(s)-1]
	}
	os.Stdout.WriteString(string(result))
}
