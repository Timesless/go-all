package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)


var pool *redis.Pool

/*
	初始化redis连接池
 */
func init() {
	pool = &redis.Pool{
		MaxActive: 16,
		MaxIdle: 4,
		IdleTimeout: time.Second * 60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.0.104:6379")
		},
	}
}

/*
	go连接redis
		方式二：go get gopkg.in/redis.v4
 */
func main() {

	// 从redis池中获取一个连接
	conn := pool.Get()
	defer conn.Close()

	// opString(conn)
	// opHash(conn)

	opList(conn)
}

/*
	操作字符串
 */
func opString(conn redis.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("执行期间发生异常...")
		}
	}()
	conn.Do("set", "k1", "v1")
	v1, _ := redis.String(conn.Do("get", "k1"))
	fmt.Println("get from redis k1 =", v1)
}

/*
	操作hash
 */
func opHash(conn redis.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("执行期间发生异常...")
		}
	}()
	conn.Do("hset", "user", "id", "001")
	conn.Do("hmset", "user", "name", "yangzl", "sex", "man")
	name, _ := redis.String(conn.Do("hget", "user", "name"))
	uinfo, _ := redis.Strings(conn.Do("hmget", "user", "name", "sex"))

	fmt.Println("get from hash name =", name)
	fmt.Printf("username = %s, sex = %s", uinfo[0], uinfo[1])
}

func opList(conn redis.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("执行期间发生异常")
		}
	}()
	conn.Do("rpush", "list1", "1", "4", "9")
	conn.Do("rpop", "list1")
	list, _ := redis.Strings(conn.Do("lrange", "list1", "0", "-1"))
	fmt.Print("get from list = ")
	for _, e := range list {
		fmt.Print(" ", e)
	}
}