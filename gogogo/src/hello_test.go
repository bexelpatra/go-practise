package hello

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// var myLogger *log.Logger

func TestHello(t *testing.T) {
	want := "Hello, world2"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}

}
func TestPanic(t *testing.T) {
	openFile("C:/goClass/iotest/test.txt")

	println("done")
}

func openFile(fn string) {
	// f, err := os.Open(fn)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// reader := bufio.NewReader(f)
	// for {
	// 	line, isPrefix, err := reader.ReadLine()
	// 	if isPrefix || err != nil {
	// 		break
	// 	}
	// b := string(line)
	// b2 := bytes.NewBuffer(line).String()
	// fmt.Println("b==b2", b == b2)
	// fmt.Println("EqualFold", strings.EqualFold(b, b2))

	// }
}

func Test_io(t *testing.T) {
	r := bufio.NewReader(os.Stdin)

	var a string
	fmt.Fscanln(r, &a)

	fmt.Println("입력받은 내용은 ", a)
}
