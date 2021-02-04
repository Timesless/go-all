package main

import (
	"fmt"
	"sync"
)

/*
	go程通信实现方式一
		以共享内存实现go程通信『传统方式，如Java, Python, C++』

	sync
		1. 互斥锁 sync.Mutex
		2. 读写锁sync.RWMutex
		sync.Once
		sync.WaitGroup
		3. 条件变量：sync.Condi + sync.Mutex
*/
func main() {
	testGoroutineSync()
}

/*
	goroutine + sync
	多个go程向同一个全局map中放入各自计算（sum）的数据
*/
func testGoroutineSync() {
	si := make([]int, 50)
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(50)

	// sum函数
	f := func(n int) int {
		rs := 0
		for i := 1; i <= n; i++ {
			rs += i
		}
		return rs
	}

	for i := 100; i < 150; i++ {
		// go程异步执行，此时i可能已经遍历了几轮，所以使用idx值copy来被捕获
		idx := i
		go func() {
			lock.Lock()
			si[idx-100] = f(idx)
			lock.Unlock()
			wg.Done()
		}()
	}
	// 阻塞，知道所有等待组执行完毕
	wg.Wait()
	for i, v := range si {
		fmt.Printf("si[%d] = %d \n", i, v)
	}
	fmt.Println("len map = ", len(si))
}
