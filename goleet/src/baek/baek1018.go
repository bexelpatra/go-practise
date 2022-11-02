package baek

import (
	"fmt"
)

var (
	board [][]string
)

func Baek1018() {
	// r := bufio.NewScanner(os.Stdin)
	// r.Scan()
	// xy := strings.Split(r.Text(), " ")
	// y, _ := strconv.Atoi(xy[0])
	// x, _ := strconv.Atoi(xy[1])
	// chass := make([][]string, y)
	board = make([][]string, 8)
	for i := 0; i < 8; i++ {
		board[i] = make([]string, 8)
		odd := i%2 == 0
		for j := 0; j < 8; j++ {
			if odd {
				if j%2 == 0 {
					board[i][j] = "W"
				} else {
					board[i][j] = "B"
				}
			} else {
				if j%2 == 0 {
					board[i][j] = "B"
				} else {
					board[i][j] = "W"
				}

			}
		}
	}

	for _, v := range board {
		fmt.Println(v)
	}

	// for i := 0; i < y; i++ {
	// 	r.Scan()
	// 	chass[i] = strings.Split(r.Text(), "")
	// }

	// for i := 0; i <= x-8; i++ {
	// 	for j := 0; j <= y-8; j++ {
	// 		check1018(&chass, i, j)
	// 	}
	// }
}

func check1018(chess *[][]string, x, y int) int {
	min := 1251
	// white
	text := "W"
	flag := false
	if text == (*chess)[y][x] {
		flag = true
	}
	count := 0
	for i := x; i < x+8; i++ {
		for j := y; j < y+8; j++ {
			if flag {
				if y%2 == 0 { // w로 시작하면서 짝수번째
					if (*chess)[x][y] == "B" {
						count += 1
					}
				} else {
					if (*chess)[x][y] == "W" {
						count += 1
					}
				}
			} else {
				if y%2 == 0 { // w로 시작하면서 짝수번째
					if (*chess)[x][y] == "W" {
						count += 1
					}
				} else {
					if (*chess)[x][y] == "B" {
						count += 1
					}
				}

			}
		}
		flag = !flag
	}
	fmt.Println(count)
	return min
}
