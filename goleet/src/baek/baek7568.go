package baek

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Baek7568() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	r.Split(bufio.ScanWords)

	defer func(){
		w.Flush()
	}()
	n := scanInt7568(r)
	people := make([][3]int, n)
	for idx := range people {
		people[idx][0] = scanInt7568(r)
		people[idx][1] = scanInt7568(r)
		people[idx][2] = idx
	}

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] && people[i][1] > people[j][1] {
			return true
		} else if people[i][0] < people[j][0] && people[i][1] < people[j][1] {
			return false
		} else if people[i][0] == people[j][0]  {
			if people[i][1] > people[j][1]{
				return true
			}else{
				return false
			}
		} else if people[i][1] == people[j][1] {
			if people[i][0] > people[j][0] {
				return true
			}else{
				return false
			}
		} else {
			return true
		}
	})
	answer := make([]int, n)
	for _, v := range people {
		fmt.Println(v)
	}
	now := 1
	plus := 0
	answer[people[0][2]] = 1

	for i := 0; i < n-1; i++ {
		j := i + 1
		if people[i][0] > people[j][0] && people[i][1] > people[j][1] {
			now += 1
			answer[people[j][2]] = now + plus
			plus =0
		} else if people[i][0] < people[j][0] && people[i][1] < people[j][1] {
			now += 1
			answer[people[j][2]] = now + plus
			plus =0
		} else if people[i][0] == people[j][0] || people[i][1] == people[j][1]{
			now += 1
			answer[people[j][2]] = now + plus
			plus =0
		} else {
			answer[people[j][2]] = now
			plus +=1
		}
	}
	fmt.Fprint(w,answer[0])
	for i := 1; i < len(answer); i++ {
		fmt.Fprint(w, " ")
		fmt.Fprint(w, answer[i])
	}

}

func Baek7568_2() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	r.Split(bufio.ScanWords)

	defer func(){
		w.Flush()
	}()
	n := scanInt7568(r)
	people := make([][3]int, n)
	for idx := range people {
		people[idx][0] = scanInt7568(r)
		people[idx][1] = scanInt7568(r)
		people[idx][2] = idx
	}

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] && people[i][1] > people[j][1] {
			return true
		} else {
			return false
		}
	})
	answer := make([]int, n)
	
	now := 1
	plus := 0
	answer[people[0][2]] = 1

	for i := 0; i < n-1; i++ {
		j := i + 1
		if people[i][0] > people[j][0] && people[i][1] > people[j][1] {
			now += 1
			answer[people[j][2]] = now + plus
			plus =0
		} else {
			answer[people[j][2]] = now
			plus +=1
		}
	}
	fmt.Fprint(w,answer[0])
	for i := 1; i < len(answer); i++ {
		fmt.Fprint(w, " ")
		fmt.Fprint(w, answer[i])
	}

}
func Baek7568_3() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	n := scanInt7568(r)
	people := make([][3]int, n)
	for idx := range people {
		people[idx][0] = scanInt7568(r)
		people[idx][1] = scanInt7568(r)
	}

	build := bytes.Buffer{}
	rank :=1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i==j {
				continue
			}
			if people[i][0] < people[j][0] && people[i][1] < people[j][1] {
				rank+=1
			}
		}
		build.WriteString(strconv.Itoa(rank))
		build.WriteString(" ")
		people[i][2]=rank
		rank=1
	}
	fmt.Println(build.String())

}
func scanInt7568(r *bufio.Scanner) int {
	r.Scan()
	num := 0
	for _, v := range r.Bytes() {
		num *= 10
		num += (int)(v - '0')
	}
	return num
}
