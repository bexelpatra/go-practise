package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Baek11650() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = scanInt11650(r)
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		} else {
			return points[i][0] < points[j][0]
		}
	})

	for _, point := range points {
		fmt.Fprintf(w, "%d %d\n", point[0], point[1])
	}
}
func scanInt11650(r *bufio.Scanner) []int {
	r.Scan()
	var a, b int
	temp := strings.Split(r.Text(), " ")
	a, _ = strconv.Atoi(temp[0])
	b, _ = strconv.Atoi(temp[1])

	point := []int{a, b}
	return point
}
