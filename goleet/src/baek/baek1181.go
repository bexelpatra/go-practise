package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Baek1181() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n int
	r.Scan()
	n, _ = strconv.Atoi(r.Text())
	// list := SortBy2{}
	list := make([]string, 0)
	for i := 0; i < n; i++ {
		r.Scan()
		list = append(list, r.Text())
	}
	// sort.Sort(list)
	sort.Slice(list, func(i, j int) bool {
		if len(list[i]) == len(list[j]) {
			return strings.Compare(list[i], list[j]) < 0
		} else {
			return len(list[i]) < len(list[j])
		}
	})

	fmt.Fprintln(w, list[0])
	for i := 1; i < len(list); i++ {
		if !strings.EqualFold(list[i-1], list[i]) {
			fmt.Fprintln(w, list[i])
		}
	}

}

// type SortBy2 []string

// func (a SortBy2) Len() int      { return len(a) }
// func (a SortBy2) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// func (a SortBy2) Less(i, j int) bool {
// 	if len(a[i]) == len(a[j]) {
// 		return strings.Compare(a[i], a[j]) < 0
// 	} else {
// 		return len(a[i]) < len(a[j])
// 	}
// }
