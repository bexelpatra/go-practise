package baek

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek1436() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	count := 0
	i := 666
	for {
		if strings.Contains(strconv.Itoa(i), "666") {
			count += 1
			fmt.Println(i)
			if count == N {
				break
			}
		}
		i += 1
	}
	fmt.Println(i)
}
