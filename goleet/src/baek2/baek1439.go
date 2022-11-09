package baek2

import (
	"bufio"
	"fmt"
	"os"
)

func Baek1439() {

	r := bufio.NewReader(os.Stdin)
	
	text,_,_ :=r.ReadLine()
	zero :=0
	one :=0
	for i := 0; i < len(text)-1; i++ {
		if text[i]==text[i+1]{
			
		}else{
			if text[i+1]=='0'{
				one+=1
			}else{
				zero+=1
			}

		}
		if i+1==len(text)-1{
				if text[i+1]=='0'{
					zero+=1	
				}else{
					one+=1
				}
			}
	}
	fmt.Println(func(a,b int)int{
		if a<b{
			return a
		}
	return b
}(zero,one))
	
}
