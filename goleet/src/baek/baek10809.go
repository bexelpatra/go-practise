package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func Baek10809() {
	r:=bufio.NewScanner(os.Stdin)
	scanText10809(r)
}

func scanText10809(r *bufio.Scanner) {
	r.Scan()
	result :=make([]int,26)
	str := bytes.Buffer{}
	for v := range result {
		result[v] = -1
	}

	for i,v := range r.Bytes() {
		if(result[v-'a']==-1){
			result[v-'a']=i
		}
	}
	for i := range result {
		str.WriteString(strconv.Itoa(result[i]))
		str.WriteString(" ")
	}

	fmt.Print(str.String())

}
