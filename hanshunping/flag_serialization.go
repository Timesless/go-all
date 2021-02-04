package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

/*
	go 命令行参数解析
		go run flag_serialization.go -u root -p 123456 --port 1521
	go 序列化与反序列化
*/
func main() {

	testFlag()
	testSerialization()
}

/*
	获取命令行参数，解析 -u, -p --port
*/
func testFlag() {
	// 命令行所有参数，以空格分割
	var args []string = os.Args
	fmt.Println(args)

	var user, password string
	var port int
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&password, "p", "", "密码")
	flag.IntVar(&port, "port", rand.Intn(65535), "端口")
	// 别忘了解析
	flag.Parse()

	fmt.Printf("-u = %s, -p = %s, --port = %d\n", user, password, port)
}

/*
	序列化测试
		1. struct序列化
		2. struct + tag 序列化
		3. slice序列化
	可以选择序列化到文件
*/
type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func testSerialization() {

	m1 := Monster{"牛魔王", 999, "男"}
	bm, _ := json.Marshal(m1)
	fmt.Println(string(bm))

	var m2 Monster
	// 忽略异常
	_ = json.Unmarshal(bm, &m2)
	fmt.Println(m2)
}
