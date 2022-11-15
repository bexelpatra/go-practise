package myTest

import (
	"fmt"
	"reflect"
)

func GoClosure() {
	funcList := make([]func(), 3)
	v := 1
	v2 := reflect.ValueOf(v).Int()
	for i := 0; i < len(funcList); i++ {
		fmt.Printf("Inserting %v \n", i)
		funcList[i] = func() {

			fmt.Println(v2)
		}
	}
	v = 2
	v2 = 4
	fmt.Printf("current funcList's v %v\n", v)

	for i := 0; i < len(funcList); i++ {
		funcList[i]()
	}

	v = 4
	fmt.Printf("current v is %v \n", v)
	for i := 0; i < len(funcList); i++ {
		funcList[i]()
	}

	v = 19999
	fmt.Printf("current v is %v \n", v)
	for i := 0; i < len(funcList); i++ {
		funcList[i]()
	}

}
