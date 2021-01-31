package main

import (
	"fmt"
	"time"
)

/*
	go
 */
func main() {

	go cyclic()
	fmt.Println("fibnacci 43 结果是: ", fibnacci(int64(43)))
	testMapReduce()

	time.Sleep(6 * time.Second)
}

/*
	打印-\|/
 */
func cyclic() {
	for {
		for _, x := range `-\|/`{
			fmt.Printf("\r%c", x)
			time.Sleep(220 * time.Millisecond)
		}
	}
}

func fibnacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibnacci(n - 1) + fibnacci(n - 2)
}

func testMapReduce() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			// 下面go程中，会阻塞等待
			time.Sleep(time.Millisecond * 500)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			tmp := <- ch1
			ch2 <- tmp * tmp
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d <- ch2\n", <- ch2)
		}
	}()
}