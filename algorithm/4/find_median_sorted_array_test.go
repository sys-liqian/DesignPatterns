package test

import (
	"fmt"
	"math"
	"testing"
)

func TestA(t *testing.T) {
	n1 := []int{2, 3, 4}
	n2 := []int{1, 6}
	findMedianSortedArrays(n1, n2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 1 {
		return float64(nums2[0])
	}
	if len(nums1) == 1 && len(nums2) == 0 {
		return float64(nums1[0])
	}
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	n1Len := len(nums1)
	n2Len := len(nums2)
	flag := (float64(n1Len) + float64(n2Len)) / 2
	intV, floatV := math.Modf(flag)
	if floatV == 0 {
		//偶数个数据,取(arr[intV]+arr[intV+1])/2
	} else {
		//奇数个数据，取arr[intV]
	}
	fmt.Println(intV)
	fmt.Println(floatV)
	return 0
}
