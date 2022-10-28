package baek

import (
	"bufio"
	"fmt"
	"os"
)

func Baek14681() {
	r:= bufio.NewScanner(os.Stdin)

	x:=isPossitive(r)
	y:=isPossitive(r)
	// if x && y{
	// 	fmt.Println(1)
	// 	return
	// }else if x &&!y{
	// 	fmt.Println(4)		
	// 	return
	// }else if !x &&y{
	// 	fmt.Println(2)
	// 	return
	// }else if !x &&!y{
	// 	fmt.Println(3)
	// 	return
	// }
	if x{
		if y{
			fmt.Println(1)
		}else{
			fmt.Println(4)
		}
	}else{
		if y{
			fmt.Println(2)
		}else{
			fmt.Println(3)			
		}
	}
}

func isPossitive(r *bufio.Scanner)bool{
	r.Scan()
	return r.Bytes()[0] !=45
}