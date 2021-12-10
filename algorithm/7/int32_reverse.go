package main

import (
	"fmt"
	"math"
	"strconv"
)

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
func main() {
	var a int32 = -123
	i := reverse(int(a))
	fmt.Println(i)
}

func reverse(x int) int {
	f := false
	if x < 0 {
		x = -x
		f = true
	}
	a := strconv.Itoa(x)
	len := len(a)
	s := ""
	for i := 0; i < len; i++ {
		s = s + fmt.Sprintf("%c", a[len-i-1])
	}
	atoi, _ := strconv.Atoi(s)
	if f {
		atoi = -atoi
	}

	if atoi < math.MinInt32 || atoi > math.MaxInt32 {
		return 0
	}
	return atoi
}
