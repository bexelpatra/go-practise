package baek

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Baek1427() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()

	n := SortBy(strings.Split(r.Text(), ""))
	sort.Sort(n)

	fmt.Println(strings.Join(n, ""))

}

type SortBy []string

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] >= a[j] }
