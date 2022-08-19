package myTest

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func nextInt2() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func Main() {
	in.Split(bufio.ScanWords)
	for t := nextInt2(); t > 0; t-- {
		fmt.Fprintln(out, main2())
	}
	out.Flush()
}

func main2() int {
	n := nextInt2()
	a := make([][2]int, n)
	for i := range a {
		a[i][0] = nextInt2()
	}
	for i := range a {
		a[i][1] = nextInt2()
	}
	var p, q, r, s int
	for i := range a {
		t, u := a[i][0], a[i][1]
		if s > q {
			t += s
		} else {
			t += q
		}
		if r > p {
			u += r
		} else {
			u += p
		}
		p, q, r, s = r, s, t, u
	}
	if r > s {
		return r
	}
	return s
}
