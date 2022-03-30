package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Baek1920() {
	// 무식하게 배열을 다 돌면서 검색 시도
	// 시간 초과
	r := bufio.NewReader(os.Stdin)
	var N, M int

	fmt.Fscanln(r, &N)
	numbers := make([]int, N)
	inputByte, _ := r.ReadBytes('\n')
	inputByte = inputByte[:len(inputByte)-2]
	input := strings.Split(string(inputByte), " ")
	for i := 0; i < N; i++ {
		numbers[i], _ = strconv.Atoi(input[i])
	}

	fmt.Fscanln(r, &M)
	has := make([]int, M)

	answer := make([]int, M)

	inputByte2, _ := r.ReadBytes('\n')
	inputByte2 = inputByte2[:len(inputByte2)-2]
	input2 := strings.Split(string(inputByte2), " ")
	for i := 0; i < M; i++ {
		has[i], _ = strconv.Atoi(input2[i])
	}

	for i, _ := range numbers {

		for j, _ := range has {
			if numbers[i] == has[j] {
				answer[j] = 1
			}
		}
	}

	for _, val := range answer {
		fmt.Println(val)
	}

}
func Baek1920_2() {
	// map을 이용한 풀이 시도
	// 시간 초과
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var N, M int

	fmt.Fscanln(r, &N)
	numbers := make(map[string]struct{}, N)
	inputByte, _ := r.ReadBytes('\n')
	inputByte = inputByte[:len(inputByte)-2]
	input := strings.Split(string(inputByte), " ")
	for i := 0; i < N; i++ {
		numbers[input[i]] = struct{}{}
	}

	fmt.Fscanln(r, &M)

	answer := make([]int, M)

	inputByte2, _ := r.ReadBytes('\n')
	inputByte2 = inputByte2[:len(inputByte2)-2]
	input2 := strings.Split(string(inputByte2), " ")

	for i, key := range input2 {
		_, y := numbers[key]
		if y {
			answer[i] = 1
		}
	}

	for _, val := range answer {
		fmt.Fprintln(w, val)
	}

}

func Baek1920_3() {
	// 이분 탐색

	r := bufio.NewReader(os.Stdin)
	var N, M int

	fmt.Fscanln(r, &N)
	numbers := make([]int, N)
	inputByte, _ := r.ReadBytes('\n')
	inputByte = inputByte[:len(inputByte)-2]
	input := strings.Split(string(inputByte), " ")
	for i := 0; i < N; i++ {
		numbers[i], _ = strconv.Atoi(input[i])
	}

	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	fmt.Fscanln(r, &M)

	inputByte2, _ := r.ReadBytes('\n')
	inputByte2 = inputByte2[:len(inputByte2)-2]
	input2 := strings.Split(string(inputByte2), " ")

	for i := 0; i < M; i++ {
		temp, _ := strconv.Atoi(input2[i])

		fmt.Println(temp)
	}
}
