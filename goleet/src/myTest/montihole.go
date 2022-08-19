package myTest

import (
	"fmt"
	"math/rand"
)

func Monti(count int, ok bool) {
	win := 0
	lose := 0

	for i := 0; i < count; i++ {

		answer := 0
		doors := make([]int, 3)
		options := []int{0, 1, 2}

		car := rand.Int() % 3
		myChoose := rand.Int() % 3

		doors[car] = 1

		if ok { // 바꾸기 전략 사용
			if car == myChoose {
				RmSlice(&options, car)
				answer = options[0] // 어떤 것을 골라도 잘못된 답
			} else {
				// mc가 보여주는 선택지에서 car제외
				// 바꾸기 전략이므로 나의 선택 제외
				// mc는 염소를 보여준다
				// 나의 선택은 바꾸기 이므로 car 선택

				RmSlice(&options, car)
				RmSlice(&options, myChoose)
				// fmt.Printf("my choise is %d, MC show %d, and i chane to another one\n", myChoose, options[0])
				answer = car
			}
		} else { // 바꾸기 전략 사용 안함
			answer = myChoose
		}

		if doors[answer] == 1 {
			win += 1
		} else {
			lose += 1
		}
	}
	fmt.Println(win, lose)
	fmt.Println(win*100/(win+lose), lose*100/(win+lose))
}

func RmSlice(s *[]int, num int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("문제발생!!!")
			fmt.Println(&s, num, len(*s))
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
