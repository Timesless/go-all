package main

import (
	"fmt"
	"sort"
)

/*
	go 接口
		接口是引用类型，默认值nil，空接口：interface{}
	golang的核心是面向接口编程
『go将接口的使用达到了极致，c++将迭代器的使用达到了极致』

	某个结构体实现接口所有方法即实现了该接口
	接口：具备某种功能，该功能可以设计为接口「is like a」
	继承：「is a」
	组合：「has a」

	接口继承「通过匿名接口字段实现，参考继承」
	接口多继承时，不能有相同的方法名

go 多态通过接口实现
	1. 多态参数「接口的方法调用传参不同调用不同的方法」
	2. 多态数组「接口数组可以存放实现该接口的所有结构体」

实践：
	Interface接口，见sort包「type Interface interface{ Len Less Swap }」
*/
func main() {

	testInterface()
	testInterfacePractice()
	testInterfacePolymorphism()
}

/*
	定义usb接口
*/
type Usb interface {
	// 函数签名：method(参数列表) 返回值列表
	Start() string
	Stop() string
}

/*
	定义结构体，并实现接口的所有方法
	IPhone, Nokia都实现了Usb接口
*/
type IPhone struct {
	name string
}

func (this *IPhone) Nice() {
	fmt.Println("nice iphone")
}

type Nokia struct {
	name string
}

/*
	实现接口的所有方法
*/
func (this *IPhone) Start() string {
	fmt.Println("iphone start")
	return this.name
}
func (this *IPhone) Stop() string {
	fmt.Println("iphone stop")
	return this.name
}

func (this *Nokia) Start() string {
	fmt.Println("nokia start")
	return this.name
}
func (this *Nokia) Stop() string {
	fmt.Println("nokia stop")
	return this.name
}

/*
	接口继承
*/
type Interface1 interface {
	test1()
}
type Interface2 interface {
	test2()
}
type Interface3 interface {
	Interface1
	Interface2
	test3()
}

/*
	接口用法
*/
func testInterface() {
	iphone, nokia := IPhone{"iphone"}, Nokia{"nokia"}
	irs := iphone.Start()
	iphone.Stop()
	nokia.Start()
	nrs := nokia.Stop()
	var u1, u2 Usb
	// 由于实现接口方法的参数是指针类型，所以iphone, nokia变量并不能赋值给usb接口变量
	// u1, u2 = iphone, nokia

	u1, u2 = &iphone, &nokia
	fmt.Println("irs =", irs, "nrs =", nrs)
	fmt.Println(u1)
	fmt.Println(u2)
}

/*
	接口实践 ===============================================
*/
type Hero struct {
	name string
	age  int
}

// 定义hero切片类型
type HeroSlice []Hero

// 排序需实现Interface接口的三个方法
// Len
// Less
// Swap
func (this HeroSlice) Len() int {
	return len(this)
}
func (this HeroSlice) Less(i, j int) bool {
	return this[i].age <= this[j].age
}

/*
	TODO
*/
func (this HeroSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func testInterfacePractice() {
	fmt.Println("========================================================")
	h1, h2, h3 := Hero{"name1", 26}, Hero{"name2", 24}, Hero{"name3", 22}
	sh := make([]Hero, 3)
	sh[0], sh[1], sh[2] = h1, h2, h3
	fmt.Println(sh)
	sort.Sort(HeroSlice(sh))
	fmt.Println(sh)
}

/*
	多态
	类型断言
*/
func testInterfacePolymorphism() {
	var usbs = make([]Usb, 2)
	usbs[0] = &IPhone{"iphone"}
	usbs[1] = &Nokia{"nokia"}
	fmt.Println(usbs)

	// 类型断言
	for _, u := range usbs {
		// 如果是iphone，那么多执行一个方法
		if ip, flag := u.(*IPhone); flag {
			ip.Nice()
		}
		u.Start()
	}
}
