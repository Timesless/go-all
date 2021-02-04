package main

/*
    go testing框架基本使用
		测试文件命名：xxx_test.go
		测试方法命名：TestXxx.go
			『测试方法X，必须不能是a-z之间的字符』

	go test：正确的测试不输出日志
	go test -v：所有测试都输出日志
*/
func main() {

}

func Add(n int) int {
	rs := 0
	for i := 0; i < n+1; i++ {
		rs += i
	}
	return rs
}

func Sub(x, y int) int {
	return x - y - 1
}
