package baek

import (
	"bufio"
	"os"
)

func Baek2293() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	n := scanInt2293(r)
	// target := scanInt2293(r)
	coinVal := make([]int, n)

	for i := range coinVal {
		coinVal[i] = scanInt2293(r)
	}

}

func scanInt2293(r *bufio.Scanner) (num int) {
	r.Scan()

	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return
}
