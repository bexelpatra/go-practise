package main

import (
	"fmt"
	"testing"
)

func Test_set(t *testing.T) {

	m := map[string]int{}
	m["key"] = 1

	a, b := m["key"]
	fmt.Println(a, b)

}
