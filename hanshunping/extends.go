package main

import (
	"fmt"
	"unsafe"
)

/*
	go 继承
		通过匿名结构体实现
	『如果是命名结构体作为属性，那么是组合的关联关系』
		父类private，public属性都能使用
		当父类与子类属性相同时，go编译器采用就近原则「单继承多继承都适用」
		当多继承A，B有相同字段时，必须显示指定使用A / B的属性
*/
func main() {
	testExtends()
	testMultiExtends()
}

type Base struct {
	name string
}

func (this *Base) Say() {
	fmt.Println("base say", this.name)
}

type Sub struct {
	Base
	name string
	age  int
}

/*func (s *Sub) Say() {
	fmt.Println("son say", s.name)
}*/

// 父类指针
type Sub2 struct {
	*Base
	name string
}

/*
	多继承 基本数据类型	Base, int
*/
type Sub3 struct {
	Base
	int
	name string
}

func testExtends() {
	b := Base{"base"}
	s := Sub{b, "sub", 1}
	fmt.Println("sizeof base =", unsafe.Sizeof(b), ", sizeof sub =", unsafe.Sizeof(s))
	fmt.Println(b)
	fmt.Println(s)

	// Say使用Base接受，打印s.name 是空串
	s.Say()
	s.Base.Say()

	s2 := Sub2{&Base{"b2"}, "s2"}
	fmt.Println(*s2.Base)
}

func testMultiExtends() {
	s3 := Sub3{Base{"b3"}, 1, "s3"}
	// 通过 instance.int来使用
	s3.int = 10
	fmt.Println(s3)
}
