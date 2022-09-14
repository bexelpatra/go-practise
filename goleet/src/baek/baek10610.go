package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"

	// "sort"
	"strings"
)

func Baek10610() {
	r := bufio.NewScanner(os.Stdin)
	str, err := scanString10610(r)
	if err != nil {
		fmt.Println(-1)
		return
	}
	sort.Slice(str, func(i, j int) bool { return str[i] > str[j] })

	fmt.Println(strings.Join(str, ""))
	// var s s
	// s := strings.Split("1238740", "")
	// s2 := "123456789"

	// sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	// sort.Slice(s2, func(i, j int) bool { return s2[i] > s2[j] })
	// fmt.Println(s)
	// fmt.Println(s2)

	// sort.Slice(s, func(i, j int) bool {
	// 	return s[i]>s[j]
	// })

	// fmt.Println(s)
	// sort.Sort(s)
	// fmt.Println(s)

}

func Baek10610_2() {
	r := bufio.NewScanner(os.Stdin)
	numbers := scanString10610_2(r)
	var buf bytes.Buffer

	if numbers[0] == 0 {
		fmt.Println(-1)
		return
	}
	var j int64
	sum := 0
	count := 10000
	for i := 9; i >= 0; i-- {
		for j = 0; j < numbers[i]; j++ {
			buf.WriteByte((byte)(i + '0'))
			if count == 0 {
				count = 10000
				sum %= 3
			}
			sum += i
			count -= 1
		}
	}
	if sum%3 == 0 {
		fmt.Println(buf.String())
	} else {
		fmt.Println(-1)
	}

}

// type s []string

// func (s s) Len() int           { return len(s) }
// func (s s) Less(i, j int) bool { return s[i] > s[j] }
// func (s s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// func scanInt10610(r *bufio.Scanner) ([]int, error) {
// 	r.Scan()
// 	bytes := r.Bytes()
// 	length := len(bytes)
// 	numbers := make([]int, length)

// 	zero := false
// 	sum := 0
// 	for i := 0; i < length; i++ {
// 		numbers[i] = int(bytes[i] - '0')
// 		sum += numbers[i]
// 		if numbers[i] == 0 {
// 			zero = true
// 		}
// 	}

// 	if !zero {
// 		return nil, fmt.Errorf("impossible!")
// 	}
// 	if sum%3 != 0 {
// 		return nil, fmt.Errorf("impossible!")
// 	}

// 	return numbers, nil
// }

func scanString10610(r *bufio.Scanner) ([]string, error) {
	r.Scan()
	bytes := r.Bytes()
	length := len(bytes)
	numbers := make([]string, length)

	zero := false
	var sum int = 0
	for i := 0; i < length; i++ {
		numbers[i] = string(bytes[i])
		sum += (int)(bytes[i] - '0')
		sum %= 3
		if bytes[i] == '0' {
			zero = true
		}
	}

	if !zero {
		return nil, fmt.Errorf("%s", "impossible!")
	}
	if sum%3 != 0 {
		return nil, fmt.Errorf("%s", "impossible!")
	}

	return numbers, nil
}

func scanString10610_2(r *bufio.Scanner) []int64 {
	result := make([]int64, 10)
	r.Scan()

	for _, v := range r.Bytes() {
		result[(int)(v-'0')] += 1
	}

	return result
}

func Baek10610_test() {
	r := bufio.NewScanner(os.Stdin)
	var by []byte
	
	r.Buffer(by,10000000)
	r.Scan()
	numbers := make([]int, 10)
	sum := 0
	for i, v := range by {
		if i % 60000== 3 {
			r.Scan()
		}
		fmt.Println(v - '0')
		numbers[v-'0'] += 1
		sum += (int)(v - '0')
	}
	if sum%3 != 0 || numbers[0] == 0 {
		fmt.Println(-1)
		return
	}
	var b bytes.Buffer
	for i := 9; i >= 0; i-- {
		for j := 0; j < numbers[i]; j++ {
			b.WriteByte((byte)(i + '0'))
		}
	}
	fmt.Println(b.String())

}
