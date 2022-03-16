package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek2606() {
	io := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	r := io.Reader
	// w := io.Writer
	var node int
	var con int
	var connection [][]int

	fmt.Fscanln(r, &node)
	fmt.Fscanln(r, &con)
	connection = make([][]int, node+1)
	for i := 0; i < con; i++ {
		var a, b int
		fmt.Fscanln(r, &a, &b)
		connection[a] = append(connection[a], b)
		connection[b] = append(connection[b], a)
	}

	q := make([]int, 0)
	q = append(q, 1)

	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		q = append(q, connection[now]...)

	}

}
