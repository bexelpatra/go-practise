package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek10799() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	s := make([]byte, 0)
	bytes := r.Bytes()
	stick := make([][]int, len(bytes)/2)
	stick[0] = make([]int, 0)
	layer := 0
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
			if bytes[idx-1] == ')' {
				// layer -= 1
				addLine10799(&stick, layer)
			}
			add10799(&stick, layer)
		}
	}

	for _, v := range stick {
		fmt.Println(v)
	}
	fmt.Println(s)
}

func pop10799(s *[]byte) bool {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp == '('
}
func addLine10799(stick *[][]int, layer int) {
	for i := layer + 1; i < len(*stick); i++ {
		if (*stick)[i] != nil {
			(*stick)[i] = append((*stick)[i], 1)
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
