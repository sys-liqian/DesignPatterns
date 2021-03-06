package algorithm

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 11, 9, 4, 10}
	//QuickSort(&arr)
	arr = InsertionSort(arr)
	fmt.Println(arr)
}

//--------------------------分割线--------------------------------
//直接插入排序

func InsertionSort(arr []int) []int {
	for i := range arr {
		pre := i - 1
		v := arr[i]
		for pre >= 0 && arr[pre] > v {
			arr[pre+1] = arr[pre]
			pre -= 1
		}
		arr[pre+1] = v
	}
	return arr
}

//--------------------------分割线--------------------------------

func QuickSort(arr *[]int) {
	quickSort(arr, 0, len(*arr)-1)
}

func quickSort(arr *[]int, start, end int) {
	if start >= end {
		return
	}
	p := partition(arr, start, end)
	quickSort(arr, start, p-1)
	quickSort(arr, p+1, end)
}

func partition(arr *[]int, start, end int) int {
	if start < end {
		pivot := (*arr)[start] //pivot保存基准值
		left := start
		right := end
		for left < right {
			for left < right && (*arr)[right] > pivot {
				right--
			}
			for left < right && (*arr)[left] <= pivot {
				left++
			}
			if left < right {
				(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
			}
		}
		(*arr)[start], (*arr)[left] = (*arr)[left], (*arr)[start]
		return left
	}
	return 0
}

//--------------------------分割线--------------------------------

//自己写的冒泡排序
func BubbleSort(src []int) []int {
	l := len(src)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if src[i] > src[j] {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
	return src
}

//也是一种冒泡排序
func BubbleSortS(src []int) []int {
	l := len(src)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
	return src
}

//--------------------------分割线--------------------------------
