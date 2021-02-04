package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/*
	go 操作fs『filesystem』
*/
func main() {

	testFileRW()
	testFileCopy()
}

/*
	文件测试
		1. 写
		2. 读
		3. copy
		4. 二进制
*/
func testFileRW() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("文件操作发生异常", err)
		}
	}()

	wfstar, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	bufWriter := bufio.NewWriter(wfstar)
	bufWriter.WriteString("hello golang\n")
	bufWriter.WriteString("hello filesystem\n")
	bufWriter.WriteString("hello filewriter\n")

	bufWriter.Flush()

	/*
		读取文件
	*/
	ofstar, rerr := os.Open("test.txt")
	if rerr != nil {
		fmt.Println("读取文件失败", rerr)
		return
	}
	bufReader := bufio.NewReader(ofstar)
	for {
		str, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}

	ofstar.Close()
	wfstar.Close()
}

/*
	一个reader使用完毕之后就没了
	要想再复制，还需要再获取一个file* 创建reader
*/
func testFileCopy() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("文件copy异常", err)
		}
	}()
	rstar, _ := os.Open("test.txt")
	reader := bufio.NewReader(rstar)

	wstar, _ := os.OpenFile("copy.txt", os.O_CREATE, os.ModePerm)
	writer := bufio.NewWriter(wstar)
	writer.ReadFrom(reader)

	rstar2, _ := os.Open("test.txt")
	reader2 := bufio.NewReader(rstar2)
	wstar2, _ := os.Create("copy2.txt")
	writer2 := bufio.NewWriter(wstar2)
	io.Copy(writer2, reader2)

	wstar2.Close()
	wstar.Close()
	rstar2.Close()
	rstar.Close()
}
