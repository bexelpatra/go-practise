package main

import (
	"fmt"
	"os"

	"example.go/goleet/baek"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// baek.Baek2606()
	// baek.Baek2606_2()
	// baek.Baek1012()
	// baek.Baek1717()
	// baek.Baek10828()
	// baek.Baek9012()
	// baek.Baek11047()
	// baek.Baek1931()
	// baek.Baek4673()
	// baek.Baek1065()
	// baek.Baek7576()
	// baek.Baek1697_2()
	// baek.Baek1026()
	// baek.Baek2798()
	// baek.Baek1541()
	// baek.Baek2217()
	// baek.Baek5585()
	// baek.Baek2231()
	// baek.Baek11724()
	// baek.Baek8958()

	// baek.Baek1152_3()
	// baek.Baek_Sample_1()
	// baek.Baek1158_2()
	// baek.Baek2750()
	// baek.Baek1427()
	// baek.Baek10989()
	// baek.Baek11650()
	// baek.Baek2309()
	// baek.Baek9465()
	// my.Do()
	// baek.Baek2579()
	// baek.Baek1003()
	// baek.Baek11053()
	// baek.Baek2164()
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i+1, " : ", baek.Sol2164_2(i+1))
	// }
	// baek.Baek2293()
	// baek.Baek2294()
	// baek.Baek1912()
	// baek.Baek12865()
	// baek.Baek12865_2()
	baek.Baek9251()
}

func gogo() {
	defer func() {
		r := recover()
		fmt.Println("00000")
		if r != nil {
			fmt.Println("error occured", r)
		}
	}()
	defer func() {
		r := recover()
		fmt.Println(11111)
		if r != nil {
			fmt.Println("error occured", r)
		}
	}()
	defer func() {

		fmt.Println(2222222)

	}()
	fi, err := os.OpenFile("d:/dddd", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		r := recover()
		fmt.Println(3333333)
		if r != nil {
			fmt.Println("error occured", r)
		}
	}()
	fi.Close()
}
