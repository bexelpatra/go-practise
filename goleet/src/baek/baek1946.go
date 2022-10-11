package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1946() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	repeat := nextInt1946(r)
	for i := 0; i < repeat; i++ {
		result := make([]int, nextInt1946(r)+1)
		for j := 0; j < len(result)-1; j++ {
			result[nextInt1946(r)] = nextInt1946(r)
		}
		do1946(result)
	}

}
func do1946(result []int) {
	num := 0
	standard := len(result) + 1
	for i := 1; i < len(result); i++ {
		if result[i] == 1 {
			fmt.Println(num + 1)
			return
		}
		if result[i] < standard {
			num += 1
			standard = result[i]
		}
	}
}
func nextInt1946(r *bufio.Scanner) int {
	num := 0
	r.Scan()
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
