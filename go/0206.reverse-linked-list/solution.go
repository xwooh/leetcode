// Created by Yuan at 2023/11/27 13:34
// leetgo: dev
// https://leetcode.cn/problems/reverse-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func insertHead(head *ListNode) (ans *ListNode) {
	// 头插法
	// 挨个把节点挪到头部

	// 具体方式:
	// 1. 从第二个节点开始，取出待插到头部的节点
	// 2. 把上一个节点的指针指向下一个节点
	// 3. 取出的节点的指针指向头结点

	p := head
	for p != nil && p.Next != nil {
		tmp := p.Next        // 取出下一个节点
		p.Next = p.Next.Next // 上一个节点指针指向下一个节点
		tmp.Next = head      // 取出节点指向头结点
		head = tmp           // 去除节点作为当前头结点
	}

	ans = head
	return
}

func reverseList(head *ListNode) (ans *ListNode) {
	ans = insertHead(head)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
