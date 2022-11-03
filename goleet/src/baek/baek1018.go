package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	board [][]string
)

func Baek1018() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	xy := strings.Split(r.Text(), " ")
	y, _ := strconv.Atoi(xy[0])
	x, _ := strconv.Atoi(xy[1])
	chass := make([][]string, y)
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

	for i := 0; i < y; i++ {
		r.Scan()
		chass[i] = strings.Split(r.Text(), "")
	}
	num := 1251
	for i := 0; i < y-8+1; i++ {
		for j := 0; j < x-8+1; j++ {
			num = min1018(check1018_2(&chass, i, j), num)
		}
	}
	fmt.Println(num)
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

func check1018_2(chass *[][]string, x, y int) int {
	a := 0
	b := 0
	for i := x; i < x+8; i++ {
		for j := y; j < y+8; j++ {
			if (*chass)[i][j] != board[i-x][j-y] {
				a += 1
			}
			if (*chass)[i][j] == board[i-x][j-y] {
				b += 1
			}
		}
	}

	return min1018(a, b)
}

func min1018(m, n int) int {
	if m > n {
		return n
	}
	return m
}
