package algorithm

import (
	"fmt"
	"testing"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func TestLengthOfLongestSubstring(t *testing.T) {
	s := "aab"
	len := lengthOfLongestSubstring(s)
	fmt.Println(len)
}

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	size := len(runes)
	if size == 1 {
		return 1
	}
	max := 0
	for j := 0; j < size; j++ {
		m := make(map[rune]int)
		for i := j; i < size; i++ {
			_, ok := m[runes[i]]
			if !ok {
				//如果元素不存在,就存进去
				m[runes[i]] = i
			} else {
				//如果存在则，输出当前map中的元素个数
				//如果字符串中无相同字母则不会进入该判断
				if len(m) > max {
					max = len(m)
				}
				break
			}
		}
		if len(m) > max {
			max = len(m)
		}
	}

	return max
}
