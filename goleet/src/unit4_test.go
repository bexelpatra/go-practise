package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type Sample_1 struct {
	A  string
	B  string
	Ab string
}
type Sample_2 struct {
	A string
	B string
}

func Test_Marshal(t *testing.T) {
	a := Sample_1{"a", "b", "c"}
	a2 := Sample_2{}
	result, _ := json.Marshal(a)
	fmt.Println(string(result))

	err := json.Unmarshal(result, &a2)
	if err != nil {
		panic("c8")
	}

	fmt.Println(print(json.Marshal(a2)))

}

func print(data []byte, a interface{}) string {
	return string(data)
}

func Test_LastIndex(t *testing.T) {
	str := "abcdefc"
	fmt.Println(strings.LastIndex(str, "c"))
}

func Test_DeleteMap(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "a"
	m["b"] = "b"
	m["c"] = "c"

	fmt.Println(m)
	delete(m, "d")
	fmt.Println(m)

}
