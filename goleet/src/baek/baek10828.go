package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek10828() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscanln(r, &n)
	s := Stack{}
	for i := 0; i < n; i++ {
		var method string
		var x int
		fmt.Fscanln(r, &method, &x)
		var result int

		switch method {
		case "pop":
			result = s.Pop()
		case "push":
			s.Push(x)
			continue
		case "size":
			result = s.Size()
		case "empty":
			result = s.Empty()
		case "top":
			result = s.Top()
		}
		fmt.Fprintf(w, "%d\n", result)
	}

}

type Stack []int

type StackInterface interface {
	Push(x int)
	Pop() int
	Size() int
	Empty() int
	Top() int
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (x int) {
	if s.Empty() == 1 {
		x = -1
		return
	}

	l := len(*s)
	x = (*s)[l-1]
	*s = (*s)[:l-1]
	return
}
func (s *Stack) Size() int {
	return len(*s)
}
func (s *Stack) Empty() (x int) {
	if s.Size() == 0 {
		x = 1
	}
	return
}
func (s *Stack) Top() (x int) {
	if s.Empty() == 1 {
		x = -1
		return
	}
	l := len(*s)
	x = (*s)[l-1]
	return

}
