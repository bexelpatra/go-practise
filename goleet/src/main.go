package main

import (
	"fmt"
	"os"
	"reflect"

	// "example.go/goleet/baek"

	"github.com/joho/godotenv"
)

func main() {
	//접기: ctrl + shift + [
	godotenv.Load(".env", ".env2")

	fmt.Println(os.Getenv("SAYMYNAME"))
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
	// baek.Baek9251()
	// baek.Baek11057()
	// baek.Baek1463_3()
	// baek.Baek1463_4()
	// baek.Baek1003_2()
	// baek.Baek9095()
	// baek.Baek11726()
	// baek.Baek11726_2()
	// baek.Baek1149()
	// baek.Baek2579()
	// baek.Baek11053()
	// baek.Baek1931()
	// baek.Baek1931_2()/
	// baek.Baek10162()/
	// baek.Baek13305()
	// baek.Baek1789()
	// baek.Baek10610()
	// baek.Baek10610_2()
	// baek.Baek10610_test()
	// myTest.Monti(11111, false)
	// baek.Baek1946()
	// baek.Baek2264()
	// baek.Baek7568()
	// baek.Baek7568_2()
	// baek.Baek7568_3()
	// baek.Baek22173()
	// baek.Baek14502()
	// baek.Baek1753()
	// baek.Baek1753_2()
	// baek.Baek1753_3()
	// baek.Hanoi()
	// baek.Baek14681()
	// baek.Baek4963()
	// baek.Baek2775()
	// baek.Baek10809()
	// baek.Baek2675()
	// baek.Baek4963()
	// baek.Baek2042()
	// baek.Baek6549()
	// baek.Baek6549_2()
	// baek.Baek1018()
	// baek.Baek1436()
	// baek.Baek1874()
	// baek.Baek4949()
	// baek.Baek10799_2()
	// baek.Baek10799_3()
	// myTest.GoClosure()
	// for {
	// baek2.Baek1406()
	// baek.Baek1717_2()
	// baek2.Baek1038()
	// baek2.Baek1038_2()
	// baek2.Baek1715()/
	// baek2.Baek1439()
	// baek2.Baek16953()
	// baek2.Baek1399()
	// baek2.Baek4796()
	// baek2.Baek1011()
	// baek2.Baek1202_3()
	// baek2.Baek1049()
	// baek2.Baek1744()
	// baek2.Baek17298()
	// baek2.Baek2864()
	// baek2.Baek1449()
	// baek2.Baek1080()
	// test2()
	// baek2.Baek11000()
	// baek2.Baek15649()
	// baek2.Baek2437_3()
	// }
}

func Test() {
	str := "1234"
	str = str[:2] + str[2+1:]

	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str[0]))
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

// func test2() {
// 	sc2.Split(bufio.ScanWords)
// 	defer wr2.Flush()

// 	N := nextInt2()
// 	arr := make([][2]int64, N<<1)
// 	for i := 0; i < N; i++ {
// 		s, t := nextInt64(), nextInt64()
// 		arr[i<<1] = [2]int64{1, s}
// 		arr[i<<1+1] = [2]int64{-1, t}
// 	}
// 	sort.Slice(arr, func(i, j int) bool {
// 		if arr[i][1] == arr[j][1] {
// 			return arr[i][0] < arr[j][0]
// 		}
// 		return arr[i][1] < arr[j][1]
// 	})

// 	concurr := int64(0)
// 	max := int64(0)
// 	for _, v := range arr {
// 		concurr += v[0]
// 		if concurr > max {
// 			max = concurr
// 		}
// 	}

// 	fmt.Fprintln(wr2, max)
// }

// //---------
// // Fast IO
// //---------

// var sc2 = bufio.NewScanner(os.Stdin)
// var wr2 = bufio.NewWriter(os.Stdout)

// func nextInt2() int {
// 	sc2.Scan()
// 	text := sc2.Text()
// 	v, _ := strconv.Atoi(text)
// 	return v
// }

// func nextInt64() int64 {
// 	sc2.Scan()
// 	text := sc2.Text()
// 	v, _ := strconv.ParseInt(text, 10, 64)
// 	return v
// }
