package main

import (
	"fmt"
	"sync"
	"time"
)

var waitNum = 0
var mutexNum = 0
var rwMutexNum = 0
var onceNum = 0
var condNum = 0
var chanNum = 0

/*

	WaitGroup「等待组」
	Once「执行一次」

	同步：「同步在不同的语义下有不同的含义，并发、数据库、缓存」
		互斥锁
		读写锁
		条件变量
		「CAS，Java中有提供」

		通道「chan，go独特的一种进程间传递数据的实现」
			Linux系统中，管道是一种进程间通信方式，两个进程借助kernel开辟的缓冲区通过管道进行数据交换
			chan 是Go设计的实现多个go程之间通信的
 */
func main() {

	go testWaitGroup()

	go testMutex()

	go testRWMutex()

	go testOnce()

	go testCond()

	go testChan()

	time.Sleep(time.Second * 2)
}

/*
	等待组
	Add
	Done
	Wait
 */
func testWaitGroup() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(num int) {
			wg.Add(1)
			waitNum ++
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("all goroutine done..., wait num =", waitNum)
}

/*
	互斥锁
	临界区保持最小颗粒度
 */
func testMutex() {
	lock := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			lock.Lock()
			mutexNum ++
			lock.Unlock()
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println("mutexNum =", mutexNum)
}

/*
	读写锁
 */
func testRWMutex()  {
	rwLock := sync.RWMutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go writer(&rwLock)
		wg.Done()
	}
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go reader(&rwLock)
		wg.Done()
	}
	wg.Wait()
}

// reader
func reader(rwLock *sync.RWMutex) {
	for {
		rwLock.RLock()
		fmt.Printf("goroutine read num = %d\n", rwMutexNum)
		time.Sleep(time.Millisecond * 500)
		rwLock.RUnlock()
	}
}
// writer
func writer(rwLock *sync.RWMutex) {
	for {
		rwLock.Lock()
		rwMutexNum ++
		fmt.Printf("goroutine write num = %d\n", rwMutexNum)
		time.Sleep(time.Millisecond * 400)
		rwLock.Unlock()
	}
}

/*
	条件变量，需配合互斥锁使用
 */
func testCond()  {
	mutex := sync.Mutex{}
	producer := sync.NewCond(&mutex)
	consumer := sync.NewCond(&mutex)

	wg := sync.WaitGroup{}
	/*
		判断
		干活
		唤醒
	 */
	for i := 0; i < 10; i++ {
		go produce(producer, &wg)
	}
	for i := 0; i < 4; i++ {
		go consume(consumer, &wg)
	}
	wg.Wait()
}

// 生产者
func produce(cond *sync.Cond, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	cond.L.Lock()
	for condNum >= 4{
		cond.Wait()
	}
	condNum ++
	fmt.Println("生产者生产了一个，剩余", condNum)
	cond.Broadcast()
	cond.L.Unlock()
}

func consume(cond *sync.Cond, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	cond.L.Lock()
	for condNum == 0 {
		cond.Wait()
	}
	condNum --
	fmt.Println("消费者消费了一个，剩余", condNum)
	cond.Broadcast()
	cond.L.Unlock()
}

/*
	确保once.Do(f func())代码只执行一次
 */
func testOnce() {
	once := sync.Once{}
	wg := sync.WaitGroup{}

	// 10个Go程只会执行一次
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				onceNum ++
			})
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("onceNum =", onceNum)
}

/*
	通道
 */
func testChan()  {
	// 带缓冲区，4个int大小
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	// 取走两个
	x, y := <- ch, <- ch
	fmt.Printf("%d = <- ch, %d = <- ch", x, y)
}
