package main

import (
	"fmt"
	"math"
	"strconv"
)

//判断回文数
func main() {
	i := 121
	palindrome := isPalindrome(i)
	fmt.Println(palindrome)
}

func isPalindrome(x int) bool {
	if x < math.MinInt32 || x > math.MaxInt32 || x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	len := len(s)
	for i := 0; i < len/2; i++ {
		if s[i] != s[len-1-i] {
			return false
		}
	}
	return true
}
