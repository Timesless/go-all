package main

import "fmt"

func init() {
	fmt.Println("初始化函数")
}

/*
    init函数
	匿名函数
	闭包：返回一个函数，该函数引用一个外部变量
		闭包本质：「是一个匿名类实例，引用 / 捕获的外部变量作为匿名类实例的成员变量」

	defer栈：释放资源，异常捕获
		涉及值的话会值拷贝，入栈

	函数调用参数传递方式
		值传递
			基本数据类型、结构体、数组
		引用传递

	Go函数不支持重载「以空接口和可变参数实现」

*/
func main() {

	testAnonymous()
	testClosure()
	testDefer(10)
}

func testAnonymous() {

	func(n1, n2 int) {
		fmt.Println("匿名函数定义时调用... n1 =", n1, ", n2 =", n2)
	}(1, 2)

	var f = func(n1, n2 int) int {
		return n1 + n2
	}
	fmt.Println("定义变量接受匿名函数n1 + n2 =", f(1, 2))
}

/*
	闭包
*/
func testClosure() {

	// 此时f为closure()调用后的结果，即func(int) int {}
	f := closure()
	f(1)
	f(1)
	ret := f(1)
	fmt.Println("闭包调用... f1 =", ret)

	f2 := closure2(10)
	f2(1)
	f2(1)
	rs := f2(1)
	fmt.Println("闭包调用... f2 =", rs)
}

/*
	closure 返回值是一个函数 func(int) int
*/
func closure() func(int) int {
	n := 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func closure2(n int) func(int) int {
	return func(i int) int {
		n += i
		return n
	}
}

func testDefer(n int) {

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer n =", n)
	n += 5
	fmt.Println("method call n =", n)
}
