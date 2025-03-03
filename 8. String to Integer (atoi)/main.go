package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
    var answer int64
    var sign int64 = 1
    started := false

    for _,l := range s{
        if l == ' '{
            if answer == 0 && started == true{
                break
            }
            continue
        }else if l == '-' && !started{
            sign = -1
            started = true
        }else if l == '+' && !started{
            started = true
            continue
        }else if l >= '0' && l <= '9'{
            answer = answer * 10 + int64(l - '0')
            started = true

//check if answer is > -2^31 and answer < 2^31-1
			if answer * sign < -2147483648{
				return 2147483648
			}else if answer * sign > 2147483647{
				return 2147483647
			}
        }else{
            break
        }
    }
    return int(answer * sign)
}