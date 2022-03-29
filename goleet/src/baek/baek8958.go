package baek

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Baek8958() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		tempBytes, _, _ := r.ReadLine()
		str := strings.Split(string(tempBytes), "")
		point := 0
		now := 0
		for i, _ := range str {
			if strings.EqualFold(str[i], "O") {
				now += 1
				point += now
			} else {
				now = 0
			}
		}
		fmt.Fprint(w, point, "\n")
	}

}
