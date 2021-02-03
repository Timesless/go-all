package main

import (
	"errors"
	"fmt"
)

/*
	go built-in
*/
func main() {

	testNew()
	testErrorHandle()
}

/*
	&n = 0xc000010090
	n2 = 0xc000010098
		由此可见，go不是通过new来分配变量在堆 / 栈上的，go有一套内存管理系统
*/
func testNew() {
	n := 100
	// n2是 *int，是一个指针
	n2 := new(int)
	*n2 = 10
	fmt.Printf("n类型 = %T, n = %d, &n = %p\n", n, n, &n)
	fmt.Printf("n2类型 = %T, n2 = %v, n2指向的值 = %d, &n2 = %p\n", n2, n2, *n2, &n2)
}

/*
	go 错误处理
		1. defer
		2. recover
		3. panic
		4. var err error = errors.New("异常")
*/
func testErrorHandle() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常被捕获...")
			// 再抛出封装后的异常
			panic(errors.New("封装后的异常"))
		}
	}()
	a, b := 10, 0
	c := a / b
	fmt.Println("a, b, c=", a, b, c)
}
