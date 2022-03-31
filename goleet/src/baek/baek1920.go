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
		idx := sort.Search(N, func(x int) bool { return numbers[x] >= temp })
		if idx < N && numbers[idx] == temp {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
func Baek1920_4() {
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
		flag := BinarySearch_2(numbers, temp)
		if flag {
			fmt.Println(1)
		} else {
			fmt.Println(0)

		}
	}
}
func Baek1920_5() {
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

	sort.Ints(numbers)

	fmt.Fscanln(r, &M)

	inputByte2, _ := r.ReadBytes('\n')
	inputByte2 = inputByte2[:len(inputByte2)-2]
	input2 := strings.Split(string(inputByte2), " ")

	for i := 0; i < M; i++ {
		temp, _ := strconv.Atoi(input2[i])
		if binarySearch_1(numbers, temp) {
			fmt.Println(1)
		} else {
			fmt.Println(0)

		}
	}
}

func Plein_de_verite() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
	}

	fmt.Fscanln(reader, &m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &b[i])
	}

	sort.Ints(a) // 정렬
	for i := 0; i < m; i++ {
		if binarySearch_1(a, b[i]) {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int {
	sc.Scan()
	text := sc.Text()
	v, _ := strconv.Atoi(text)
	return v
}
func Baek_Sample_1() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n1 := scanInt()
	a := make([]int, n1)
	for i := range a {
		a[i] = scanInt()
	}
	n2 := scanInt()
	b := make([]int, n2)
	for i := range b {
		b[i] = scanInt()
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	for _, bv := range b {
		idx := sort.SearchInts(a, bv)
		if idx < n1 && a[idx] == bv {
			wr.WriteString("1\n")
		} else {
			wr.WriteString("0\n")
		}
	}
}

func binarySearch_1(array []int, val int) bool {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if array[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(array) || array[low] != val {
		return false
	}
	return true
}
func BinarySearch_2(list []int, val int) (flag bool) {
	start := 0
	end := len(list) - 1

	for start <= end {
		mid := (start + end) / 2
		if list[mid] < val {
			start = mid + 1
		} else if list[mid] > val {
			end = mid - 1
		} else {
			return true
		}
	}
	return
}
