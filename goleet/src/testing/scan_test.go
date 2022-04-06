package testing

import (
	"bufio"
	"os"
)

func Do() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
}

func nextInt(in *bufio.Scanner) int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}
