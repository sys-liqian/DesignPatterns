package main

import (
	"fmt"
	"math/rand"
)

/*
golang 内存逃逸分析

go run -gcflags "-m -m" main.go
可以产看所有编译器优化
go run -gcflags "-m -m -l"
禁用内联优化，只关注逃逸优化
*/
func main() {

}

/*
内存逃逸情况1：返回局部变量指针
.\main.go:24:2: result escapes to heap:
.\main.go:24:2:   flow: ~r2 = &result:
.\main.go:24:2:     from &result (address-of) at .\main.go:25:9
.\main.go:24:2:     from return &result (return) at .\main.go:25:2
.\main.go:24:2: moved to heap: result
*/
func Add(a, b int) *int {
	result := a + b
	return &result
}

/*
内存逃逸情况2：无法确定变量类型（interface）
*/
func CallInterface() {
	str := "aaa"
	//调用该行函数，由于fmt.Printf入参是interface类型，所以str发生了逃逸
	//str escapes to heap只是str存储的值逃逸到了堆上，但是str变量并没有在堆上分配
	fmt.Printf("%v", str)
	//调用该行函数,由于访问了str的地址，&str被包装成了interface传入fmt.Printf
	//所以&str也需要在堆上分配，由于堆上不能存储一个栈上的地址，所以str变量也逃逸到了堆上
	//即moved to heap
	fmt.Printf("%p", &str)
}

/*
内存逃逸情况3：变量大小不确定或者栈空间不足64kb
*/
func LessThan8192() {
	//int64 8byte 128个int为1kb 8192个int为64k
	nums := make([]int, 8192)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Int()
	}
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
