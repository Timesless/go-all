package main

import "testing"

/*
	测试
*/

func TestAdd(t *testing.T) {
	r1 := Add(10)
	if 55 != r1 {
		t.Fatal("Add执行异常, 期望 =", 55, "实际结果 =", r1)
	} else {
		t.Log("结果符合预期")
	}
}

func TestSub(t *testing.T) {
	r := Sub(5, 1)
	if r != 4 {
		t.Fatal("sub执行异常, 期望 =", 4, "实际结果 =", r)
	} else {
		t.Log("结果符合预期")
	}
}
