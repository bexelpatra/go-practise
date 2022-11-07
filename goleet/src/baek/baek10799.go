package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek10799() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()

	bytes := r.Bytes()
	stick := make([][]int, len(bytes)*2)
	stick[0] = make([]int, 0)
	layer := 0
	beforeLayer := 0
	for idx, v := range bytes {
		if v == '(' {
			layer += 1
		} else if v == ')' {
			layer -= 1
			for i := 0; i <= layer; i++ {
				if stick[i] == nil {
					stick[i] = make([]int, 1)
				}
			}
			if bytes[idx-1] == '(' {
				add10799(&stick, layer)
			} else {
				if beforeLayer > layer {
					addLine10799(&stick, layer)
				}
			}
		}
		beforeLayer = layer
	}
	result := 0
	for _, v := range stick {
		fmt.Println(v)
		for _, v2 := range v {
			if v2 != 0 {
				result += (v2 + 1)
			}
		}
	}
	fmt.Println(result)
}

func addLine10799(stick *[][]int, layer int) {
	for i := layer + 1; i < len(*stick); i++ {
		if (*stick)[i] != nil {
			(*stick)[i] = append((*stick)[i], 0)
		}

	}
}

func add10799(stick *[][]int, layer int) {
	for i := 1; i < layer+1; i++ {
		if (*stick)[i] != nil {
			(*stick)[i][len((*stick)[i])-1] += 1
		}
	}
}

func Baek10799_2() {
	r := bufio.NewScanner(os.Stdin)
	r.Buffer(make([]byte, 100100), 100100)
	r.Scan()

	text := r.Text()
	result := 0
	stack := make([]rune, 0)
	prev := ' '
	for _, v := range text {
		if v == '(' {
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1]
			if prev == '(' {
				result += len(stack)
			} else {
				result += 1
			}
		}
		prev = v
	}

	fmt.Println(result)
}
func Baek10799_3() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()

	layer := 0
	result := 0

	text := r.Text()
	prev := ' '

	for _, v := range text {
		if v == '(' {
			layer += 1
		} else {
			layer -= 1
			if prev == '(' {
				result += layer
			} else {
				result += 1
			}
		}
		prev = v
	}

	fmt.Println(result)
}
func Baek10799_4() {

}
