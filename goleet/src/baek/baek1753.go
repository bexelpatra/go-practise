package baek

import (
	"bufio"
	"os"
)

var (
	V int
	E int
)

func Baek1753() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	V = scanInt1753(r)
	E = scanInt1753(r)
	// root := scanInt1753(r)
	graph := make([][]int, V)
	for i := range graph {
		graph[i] = make([]int, 2)
	}
	for i := 0; i < E; i++ {
		a := scanInt1753(r)
		b := scanInt1753(r)
		c := scanInt1753(r)

		if graph[a][b] > c {

		}
		graph[a][b] = c

	}

}

func scanInt1753(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
