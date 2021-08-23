package main 

import "fmt"
import . "math"

func reverse(x int) int {
    var isNegatif bool
    if x < 0 {
        isNegatif = true
        x = -x
    }
    
    var res int
    for x > 0 {
        res = res * 10 + x % 10
        x = x / 10
    }
    
	maxInt32 := int32(MaxInt32)
    if res > int(maxInt32) {
        return 0
    }
    
    if isNegatif != true{
        return res
    } else {
        return -res
    }
}

func main() {
	res := reverse(1534236469) // 9646324351
	fmt.Println(res)
}