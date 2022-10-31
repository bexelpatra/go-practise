package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek2675() {
	r:= bufio.NewScanner(os.Stdin)
	N,_ :=strconv.Atoi(scanText2675(r))
	w :=bytes.Buffer{}
	for i := 0; i < N; i++ {
		arr := strings.Split(scanText2675(r), " ")
		loop,_:= strconv.Atoi(arr[0])
		for _,v := range strings.Split(arr[1], "") {
			for j := 0; j <loop ; j++ {
				w.WriteString(v)	
			}
		}
		w.WriteString("\n")
	}
	fmt.Print(w.String())

}
func scanText2675(r *bufio.Scanner) string{
	r.Scan()
	return (r.Text())
}