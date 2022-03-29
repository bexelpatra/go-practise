package baek

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Baek1152() {
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	fmt.Println(strings.Replace(input, "\n", "", 2))
	fmt.Println(len(strings.Split(strings.Trim(input, " "), " ")))
	// count := 0
	// for _, val := range strings.Split(input, " ") {

	// 	if len(val) > 0 && val != "\n" {
	// 		count += 1
	// 	}
	// }
	// fmt.Println(count)
}
func Baek1152_2() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	words := strings.Split(input, " ")
	var count int
	for i := range words {
		if words[i] != "\n" && words[i] != "" {
			count++
		}
	}
	fmt.Println(count)
}

func Baek1152_3() {
	// readline의 경우 읽을수 있는 크기의 한계가 있다.
	// 해결방법
	// 1. readline 이용
	// 2. readbyte 이용
	// 3. readline 의 prefix값 이용?

	r := bufio.NewReader(os.Stdin)

	input, prefix, _ := r.ReadLine()
	for prefix {
		var temp []byte
		temp, prefix, _ = r.ReadLine()
		input = append(input, temp...)
	}
	words := strings.Split(string(input), " ")

	count := 0
	for i := range words {
		if words[i] != "" && words[i] != "\n" {
			count += 1
		}
	}
	fmt.Println(count)

}
func Baek1152_4() {
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadBytes('\n')
	words := strings.Split(string(input), " ")

	count := 0
	for i := range words {
		if words[i] != "" && words[i] != "\n" {
			count += 1
		}
	}
	fmt.Println(count)
}
