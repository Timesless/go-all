package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
   go 字符串API
*/
func main() {
	testString()
	testStringSlice()
}

func testString() {

	var s1, s2 string = "hello, go", "111"
	i1 := strings.Compare(s1, "hello, goo")
	i2, _ := strconv.Atoi(s2)
	fmt.Println("i1 =", i1, ", i2 =", i2)

	b1 := strings.EqualFold(s1, "HELLO, go")
	fmt.Println("flag b =", b1)

	b2, b3 := strings.HasPrefix(s1, "He"), strings.HasSuffix(s1, "go")
	fmt.Println("hello world has prefix He ?", b2, " has suffix ld ?", b3)

	s3 := "golang你好"
	var r1 = []rune(s3)
	fmt.Println("golang 你好 长度 =", len(s3), "[]rune 长度 =", len(r1))
	for i, e := range r1 {
		fmt.Printf("s3[%d] = %c ", i, e)
	}
	i3, i4, i5 := strings.IndexByte(s3, 'H'), strings.IndexRune(s3, '你'), strings.IndexAny(s3, "cdg")
	fmt.Printf("\ni3 = %d, i4 = %d, i5 = %d \n", i3, i4, i5)

	var f func(rune) bool = func(r rune) bool {
		return r > 'l'
	}
	i6 := strings.IndexFunc(s1, f)
	fmt.Println("s1 index func =", i6)
}

func testStringSlice() {
	s1, s2 := "hello go", "golang 你好"
	bs, rs := []byte(s1), []rune(s2)
	bs[0] = 'z'
	rs[7] = '您'
	s1, s2 = string(bs), string(rs)
	fmt.Println("s1 =", s1, ", s2 =", s2)
}
