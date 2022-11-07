package baek2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Baek1406() {
	r := bufio.NewReader(os.Stdin)

	tempB, _, _ := r.ReadLine()
	s := string(tempB) + " "
	str1406 := Str1406{[]rune(s), len(s) - 1}

	loopStr, _, _ := r.ReadLine()
	loop, _ := strconv.Atoi(string(loopStr))

	for i := 0; i < loop; i++ {
		text, _, _ := r.ReadLine()
		switch text[0] {
		case 'L':
			temp := str1406.Cursor
			str1406.Cursor = str1406.Left1406()
			str1406.Swap1406(temp, str1406.Cursor)
		case 'D':
			temp := str1406.Cursor
			str1406.Cursor = str1406.Right1406()
			str1406.Swap1406(temp, str1406.Cursor)
		case 'B':
			str1406.B1406()
		case 'P':
			str1406.P1406(string(text[2:]))
		}
	}
	fmt.Println(str1406.Print())
}

type Str1406 struct {
	Str    []rune
	Cursor int
}

func (s *Str1406) Swap1406(a, b int) {
	s.Str[a], s.Str[b] = s.Str[b], s.Str[a]
}
func (s *Str1406) Left1406() int {
	temp := s.Cursor
	if s.Cursor > 0 {
		temp -= 1
	}
	return temp
}
func (s *Str1406) Right1406() int {
	temp := s.Cursor
	if len(s.Str)-1 > s.Cursor {
		temp += 1
	}
	return temp
}
func (s *Str1406) B1406() {
	left := s.Str[:s.Left1406()]
	s.Str = append(left, s.Str[s.Cursor:]...)
	s.Cursor = s.Left1406()
}
func (s *Str1406) P1406(add string) {
	l := len(add)
	if s.Cursor != len(s.Str)-1 {
		add = add + " "
	}
	cursor := s.Cursor
	left := s.Str[:cursor]
	temp := append([]rune(add), s.Str[s.Right1406():]...)
	s.Str = append(left, temp...)
	s.Cursor = s.Cursor + l
}
func (s *Str1406) Print() string {
	var temp []rune
	if s.Right1406() == s.Cursor {
		temp = s.Str[:s.Cursor]
	} else {
		temp = append(s.Str[:s.Cursor], s.Str[s.Right1406():]...)
	}
	return string(temp)
}
