package main

import (
	"fmt"
	"log"
	"time"
)

/*

go 提供了两种并发形式
	1. 多线程并发模型
	goroutine + sync
		通过共享内存的方式进行线程间通信
			例如：在访问共享数据时通过锁来访问
	2. CSP『communicating sequential process』并发模型
	goroutine + chan
		CSP: “以通信的方式来共享内存”
			goroutine + channel『并发结构体(goroutine)之间的通信机制』

go程「协程」基本使用
	1. 有独立的栈空间
	2. 共享堆空间
	3. 调度由用户控制
	4. 协程是轻量级线程「线程是轻量级进程」

	内核级线程(KLT, 由操作系统提供线程实现『Java, C++』)
		OS控制生命周期，由OS提供执行栈
	用户级线程(ULT, 操作系统线程：用户线程 = 1：N)
		用户控制生命周期，由语言库提供执行栈

	MPG
*/
func main() {

	goroutineException()
	testGoroutine()
}

func testGoroutine() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("hello goroutine")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("hello golang")
		time.Sleep(100 * time.Millisecond)
	}
}

/*
	go程异常处理
*/
func goroutineException() {

	defer func() {
		if e := recover(); e != nil {
			log.Fatal("go 程执行异常", e)
		}
	}()
	rs := 1 / 0
	fmt.Println(rs)
}
