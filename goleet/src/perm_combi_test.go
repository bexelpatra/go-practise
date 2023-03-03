package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	r1 := 3
	temp1 := make([]int, r1)
	perm_unduplicapable(&list, &temp1, 3, 0, 0)
	// fmt.Println(temp1)
	r2 := 3
	temp2 := make([]int, r2)
	perm_duplicapable(&list, &temp2, 3, 0, 0)
	// fmt.Println(temp2)
}
func perm_unduplicapable(list *[]int, temp *[]int, n, current, start int) {
	if n == current {
		fmt.Println("복사안됨", *temp)
	} else {
		for i := start; i < len(*list); i++ {
			(*temp)[current] = (*list)[i]
			perm_unduplicapable(list, temp, n, current+1, start+1)
		}
	}
}

func perm_duplicapable(list *[]int, temp *[]int, n, current, start int) {
	if n == current {
		fmt.Println("복사가능", *temp)
		return

	} else {
		for i := start; i < len(*list); i++ {
			(*temp)[current] = (*list)[i]
			perm_duplicapable(list, temp, n, current+1, start)
		}
	}
}
