package design

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []string{"1"}
	fmt.Printf("主函数中slice地址：%p\n", slice)
	fmt.Printf("主函数slice数组地址：%p\n", &slice[0])
	add(&slice, "2")
	fmt.Printf("add后主函数中slice地址：%p\n", slice)
	fmt.Printf("add后主函数slice数组地址：%p\n", &(slice[0]))
	fmt.Println(slice)
}

func add(slice *[]string, n string) {
	fmt.Printf("add 函数中slice地址：%p\n", slice)
	fmt.Printf("add slice数组地址：%p\n", &((*slice)[0]))
	*slice = append(*slice, n)
	fmt.Printf("add函数append后slice地址：%p\n", slice)
}

func TestConvert(t *testing.T) {
	var a = 1
	f := float64(a)
	fmt.Printf("%f", f)
}
