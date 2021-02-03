package main

import (
	"fmt"
	"time"
)

/*
    时间日期函数API
	Format格式固定「2006 1 2 3 4 5」
		2006/01/02 15:04:05
*/
func main() {

	testDateTime()
}

func testDateTime() {
	time.Sleep(time.Millisecond * 1)
	fmt.Println("time.Now =", time.Now())
	fmt.Println("格式化日期 now =", time.Now().Format("2006-01-02 15:04:05"))

	var now time.Time = time.Now()
	fmt.Println("unix time =", now.Unix(), ", unixNano time=", now.UnixNano())
}
