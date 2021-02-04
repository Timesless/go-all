package main

import (
	"fmt"
	"time"
)

/*
	go程通信实现二『CSP』
	“以通信的方式来共享内存”

	管道是队列结构，无法通过索引访问，只能访问队头队尾，所以for遍历只返回一个值
	chan 关闭
	chan 遍历

*/
func main() {
	// testGoroutineChan()
	testChanSelect()
}

/*
需求：
	1. 一个go程向一个chan中写入数据，另一个go程读取数据
	2. 当数据读取完毕之后，结束程序
	3. 请使用goroutine + chan完成
*/
func testGoroutineChan() {

	// 定义一个读写数据的chan
	// 一个退出的标记chan
	// main线程，读取exitChan，直到dataChan被读取完毕后向exitChan写入数据并关闭，main线程才得以执行退出
	dataChan, exitChan := make(chan int, 20), make(chan bool, 1)

	// writer
	go func() {
		for i := 0; i < 20; i++ {
			dataChan <- i
			fmt.Println("write data", i)
			time.Sleep(20 * time.Millisecond)
		}
		close(dataChan)
	}()

	// reader
	go func() {
		for {
			i, ok := <-dataChan
			if !ok {
				break
			}
			fmt.Println("read data ", i)
			time.Sleep(100 * time.Millisecond)
		}
		// break之后向exitChan中写入数据并关闭
		exitChan <- true
		close(exitChan)
	}()

	// 阻塞读取exitChan
	<-exitChan
	fmt.Println("mission complete")
}

/*
	可以声明只读 / 只写的chan

	通常的做法：
		声明一个可读可写的chan，传递给方法时定义为 只读 / 只写
*/
func testRWChan() {

	// send only chan
	var rc = make(chan<- int, 3)
	// recieve only chan
	var wc = make(<-chan int, 3)

	rc <- 3
	// num := <-rc

	//wc <- 3
	num := <-wc
	fmt.Println(num)
}

/*
	select chan
	在某些情况，无法确定何时关闭管道时，通过select管道也可以退出
*/
func testChanSelect() {

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

outer:
	for {
		select {
		case v := <-ch:
			fmt.Println("读取数据 =", v)
		default:
			close(ch)
			fmt.Println("chan读取不到数据了，关闭chan")
			break outer
		}
	}
}
