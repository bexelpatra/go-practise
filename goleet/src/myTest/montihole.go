package myTest

import (
	"fmt"
	"math/rand"
)

func Monti(count int) {
	win := 0
	lose := 0

	for i := 0; i < count; i++ {

		answer := 0
		doors := make([]int, 3)
		chooses := []int{0, 1, 2}

		car := rand.Intn(3)
		myChoose := rand.Intn(3)

		doors[car] = 1
		if car == myChoose {
			RmSlice(&chooses, car)
			answer = chooses[rand.Intn(2)%2]
		} else {
			RmSlice(&chooses, car)
			RmSlice(&chooses, myChoose)
			answer = chooses[0]
		}
		if doors[answer] == 1 {
			win += 1
		} else {
			lose += 1
		}
	}
	fmt.Println(win, lose)
}

func RmSlice(s *[]int, num int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("문제발생!!!")
			fmt.Println(&s, num)
		}

	}()
	for i, v := range *s {
		if v == num {
			(*s)[i] = (*s)[len(*s)-1]
			(*s) = (*s)[:len(*s)-1]
			return
		}
	}
}
