package baek

import (
	"bufio"
)

// 가방챙기기
// 점화식 세우기가 어렵다...
func Baek12865() {
	// r := bufio.NewScanner(os.Stdin)
	// r.Split(bufio.ScanWords)
	// maxItem := nextInt12865(r)
	// maxWeight := nextInt12865(r)

	// list := make([][]int, maxItem)
	// for i := 0; i < maxItem; i++ {
	// 	list[i] = []int{nextInt12865(r), nextInt12865(r)}
	// }

}

func nextInt12865(r *bufio.Scanner) int {
	num := 0
	r.Scan()
	for _, val := range r.Bytes() {
		num *= 10
		num += int(val - '0')
	}
	return num
}
