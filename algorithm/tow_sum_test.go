package algorithm

import (
	"fmt"
	"testing"
)

/**
给定一个整数数组 num 和一个整数目标值 target，请你在该数组中找出 和为目标值 target的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
*/
func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	sum := twoSum1(nums, target)
	fmt.Println(sum)
}

//典型的以时间换空间,时间复杂度O(n*n)
func twoSum(nums []int, target int) []int {
	for k1, v1 := range nums {
		for i := k1 + 1; i < len(nums); i++ {
			if v1+nums[i] == target {
				return []int{k1, i}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k1, v1 := range nums {
		if v2, ok := m[target-v1]; ok {
			return []int{v2, k1}
		} else {
			m[v1] = k1
		}
	}
	return []int{}
}
