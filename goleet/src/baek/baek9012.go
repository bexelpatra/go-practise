package baek

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Baek9012() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		s := Stack{}
		var line string
		fmt.Fscanln(r, &line)
		answer := "YES"
	strLoop:
		for _, str := range line {
			if strings.EqualFold(string(str), "(") {
				s.Push(9)
			} else {
				pop := s.Pop()
				if pop != 9 {
					answer = "NO"
					break strLoop
				}
			}
		} // strLoop end
		if s.Size() > 0 {
			answer = "NO"
		}
		fmt.Fprintf(w, "%s\n", answer)
	}
}
