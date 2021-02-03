package main

import "fmt"

/*
	字段小写为private，只能在本包内访问
	字段大写为public
	结构体的字段是连续分配的，且会存在对齐处理

	结构体是值类型
	结构体的指针可以直接 dot 属性，go做了处理「c/cpp中使用->而不是dot，标准写法是(*instance).field」

	golang 结构体无构造函数「使用工厂模式提供实例构造方法」
*/
type Cat struct {
	name, color string
	age         int
}

/*
    go struct
		go 不是纯粹的OOP
	go 没有方法重载、构造函数、析构函数、this指针
	go 继承通过匿名字段实现
*/
func main() {
	testStruct()
}

func testStruct() {
	var c1 Cat
	c1.name = "小白"
	// cp, cp2都是*Cat
	var cp *Cat = &Cat{"小黑", "黑色", 2}
	cp2 := new(Cat)
	c2 := Cat{
		name:  "小花",
		color: "rainbow",
		age:   1,
	}
	fmt.Println("小白 =", c1, "小花 =", c2)
	// (*cp).age == cp.age
	(*cp).age = 1
	fmt.Println("*cp =", *cp, "*cp2 =", *cp2)
}

/*
	结构体内存分配机制
*/
func testStructMemroyAllocate() {

}
