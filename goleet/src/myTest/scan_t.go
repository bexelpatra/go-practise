package myTest

import (
	"bufio"
	"fmt"
	"os"
)

func Do() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	loop := nextInt(r)
	list := make([]int, 0)
	for i := 0; i < loop; i++ {
		list = append(list, nextInt(r))
	}
	fmt.Printf("%v\n", list)
}

func nextInt(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, c := range r.Bytes() {
		num *= 10
		num += int(c - '0')
	}
	return num
}

func UnconditionedLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
