package main

import "fmt"

type Dog struct {
	name, color string
	age         int
}

/*
	go 自定义类型的方法

	golang 将结构体和结构体指针设计为同一API使用

	推荐在实例方法上使用指针（前提是这个类型不是一个自定义的 map、slice 等引用类型）
	当结构体较大的时候使用指针会更高效
	如果要修改结构内部的数据或状态必须使用指针
	当结构类型包含 sync.Mutex 或者同步这种字段时，必须使用指针以避免成员拷贝
	如果你不知道该不该使用指针，使用指针
*/
func main() {
	testMethod()
	testFactoryModel()
}

// 方法接受指针「无论调用时传递的是值还是指针」
func (this *Dog) SayHello() {
	this.name = "zzz"
	fmt.Println(this.name + " say hello")
}

// 方法接受值「无论实参是值还是指针，都拷贝值」
func (this Dog) SayHello2() {
	this.name = "sss"
	fmt.Println(this.name + " say hello")
}

// 重写String，打印时会自动调用
func (this *Dog) String() string {
	return fmt.Sprintf("dod name = %s, color = %s, age = %d", this.name, this.color, this.age)
}

/*
	工厂模式 提供实例构造方法
		可以返回类型本身
		或者返回指针
*/
func NewDog(name, color string, age int) *Dog {
	return &Dog{name, color, age}
}

func testMethod() {
	d1 := Dog{"tom", "blue", 2}
	d1.SayHello()
	fmt.Println(d1.name)
	(&d1).SayHello()
	fmt.Println(d1.name)
	d1.SayHello2()
	fmt.Println(d1.name)
	(&d1).SayHello2()
	fmt.Println(d1.name)

	fmt.Println(d1)
}

/*
	工厂模式构造实例「当结构体名称小写时提供」
*/
func testFactoryModel() {
	d := NewDog("jerry", "yellow", 2)
	// 调用重写的String
	fmt.Println(&*d)
	// fmt.Println(*d)
}
