package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	val int
	next *Node
}

type Deque struct {
	// 首尾节点引用
	head, tail *Node
	// 双端队列容量
	size int
}

// 构造器
func NewNode(val int, node *Node) *Node {
	return &Node{val, node}
}

// 尾插
func AddLast(head *Node, val int) {
	node := NewNode(val, nil)
	tmp := head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = node
}

// 头插
func AddFirst(head *Node, val int) {
	node := NewNode(val, head.next)
	head.next = node
}

func PeekFirst(head *Node) int {
	return head.next.val
}

func PeekLast(head *Node) int {
	return 0
}

// 查看队头

func ToString(head *Node) string {
	str := ""
	tmp := head.next
	for tmp.next != nil {
		str += strconv.Itoa(tmp.val) + " -> "
		tmp = tmp.next
	}
	if tmp != nil {
		str += strconv.Itoa(tmp.val)
	} else {
		return "empty queue"
	}
	return str
}


/*
	链表实现双端队列
	TODO 暂未实现
 */
func main() {

	head := NewNode(-1, nil)

	AddLast(head, 4)
	AddLast(head, 5)
	AddFirst(head, 1)
	AddFirst(head, 2)
	fmt.Println(ToString(head))

}
