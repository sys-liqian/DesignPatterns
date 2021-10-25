package a20

import (
	"fmt"
	"strings"
	"testing"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

*/
func Test(t *testing.T) {
	str1 := "([])"
	result := isValid(str1)
	fmt.Println(result)
}

func isValid(s string) bool {
	len := len(s)
	//如果字符个数不为偶数,则肯定不能闭合
	if len%2 != 0 {
		return false
	}
	len = len / 2
	for i := 0; i < len; i++ {
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)
	}
	return s == ""

}

/**
还可以用栈实现

*/
