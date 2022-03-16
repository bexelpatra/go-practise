package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"
)

type M map[string]string
type A int

func Test_1(t *testing.T) {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s \n", bytes.Join(s, []byte{}))

	a := []byte("abc")

	fmt.Println(a)
	fmt.Printf("%s\n", a)

	var jido M = make(M)
	jido["a"] = "v"
	jido["45"] = "hvn"
	fmt.Println(jido)

	block := NewBlock("a", []byte("asd"))
	fmt.Println(block)
}

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func Test_io(t *testing.T) {
	var n int
	fmt.Fscan(in, &n)
	fmt.Println(n)

	defer out.Flush()
}

func Test_ioEfficency(t *testing.T) {
	a := make([]int, 100_000)

	start := time.Now()
	for i := 0; i < 100_000; i += 1 {
		fmt.Scan(&a[i])
	}
	fmt.Printf("fmt.Scan : %v\n", time.Since(start))

	start = time.Now()
	r := bufio.NewReader(os.Stdin)

	for i := 0; i < 100_000; i += 1 {
		fmt.Fscan(r, &a[i])
	}
	fmt.Printf("fmt.FScan : %v\n", time.Since(start))

}

func Test_bigInt(t *testing.T) {
	a := big.NewInt(10).Add(big.NewInt(2), big.NewInt(3))

	fmt.Println(a)
}

func Test_swap(t *testing.T) {

	data := []byte("simple test")
	a, b := 1, 2

	a, b = b, a
	w := bufio.NewWriter(os.Stdin)
	fmt.Fprintln(w, a, b)
	w.Flush()

	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	// fmt.Println(string(data))

	for i, b := range data {
		fmt.Println(i, b, string(b))
	}
}

func Test_dfs(t *testing.T) {
	ints := make([]int, 1)
	ints = append(ints, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ints = append(ints[:2], ints[3:]...)
	for _, num := range ints {
		fmt.Println(num)
	}
}
