package main

import (
	"fmt"
	"math/rand"
	"time"
)

var global int

/*
	数组与切片
		数组：
		1. 长度也是数组类型的一部分
		2. 不能像C++一样通过指针访问数组
		3. 数组每个元素有默认零值
		4. 作为参数传递时为值传递
		5. fori，for range遍历
*/
func main() {
	testArr()

	testSlice()

	testSliceAppend()
}

func testArr() {
	a1 := [3]int{1, 2, 3}
	var a2 = [3]int{2, 3, 4}
	var a3 = [...]int{4, 5, 6}

	fmt.Println("a1 =", a1, "a2 =", a2, "a3 =", a3)

	for i, ln := 0, len(a1); i < ln; i++ {
		fmt.Print(a1[i])
	}
	fmt.Println()
	for _, e := range a1 {
		fmt.Print(e)
	}
	fmt.Println()

	testArrParam(&a1, a2)
	fmt.Println("a1 =", a1, "a2 =", a2)

	practiseArr()
}

/*
	arr是指针
	arr2传递的是值
*/
func testArrParam(arr *[3]int, arr2 [3]int) {
	// (*arr)[0] == arr[0]
	(*arr)[0] = 10
	arr2[0] = 10
}

/*
	数组实践
*/
func practiseArr() {
	var arr [5]int
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	for i, ln := 0, len(arr); i < ln; i++ {
		arr[i] = rand.Intn(10)
	}
	fmt.Println(arr)

	for i, ln := 0, len(arr); i < ln/2; i++ {
		tmp, r := arr[i], ln-1-i
		arr[i] = arr[r]
		arr[r] = tmp
	}
	fmt.Println(arr)
}

/*
	切片
		由打印得出
		1. s1, s2, s3 在堆栈上分配「毋庸置疑」
		2. s1, s2, s3的元素都是在栈上分配的
		3. 栈在高地址，从高地址向低地址增长。
		4. 堆在低地址，从低地址向高地址增长。
		5. go内存分配是自己管理的
*/
func testSlice() {
	var s1 []int
	s1 = make([]int, 3, 6)
	s1[0] = 10
	heapInt := new(int)
	s1[1] = *heapInt
	s1 = append(s1, 10)
	var s2 = []int{1, 2, 3}

	arr := [...]int{1, 2, 3, 4, 5}
	s3 := arr[1:4]
	// 切片引用数组所有元素
	s3 = arr[:]
	s3 = arr[1:]
	s3 = arr[:4]

	fmt.Println("s1 =", s1, ", s2.len =", len(s2), ", s2.cap =", cap(s2), ", s2 =", s2, ", s3 =", s3)
	fmt.Println("global* =", &global)
	fmt.Printf("s1* = %p, s1[0]* = %p, heapInt = s1[1]* = %p, %p,\narr* = %p,arr[1]* = %p, s3* = %p, s3[0]* = %p",
		&s1, &s1[0], heapInt, &s1[1], &arr, &arr[1], &s3, &s3[0])
}

func testSliceAppend() {
	s1 := make([]int, 3, 5)
	s1[0], s1[1] = 10, 10

	// 创建新的数组，将数组元素值copy过去
	s2 := append(s1, 10, 10)
	s2 = append(s2, s1...)
	s2 = append(s2, 1, 1)
	s2[0] = 100

	// 深拷贝
	copy(s1, s2)

	// &s1[0] != &s2[0]
	fmt.Println("\n&s1[0] =", &s1[0], ", &s2[0] =", &s2[0])
	fmt.Println(s1)
	fmt.Println(s2)

}
