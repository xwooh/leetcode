// Created by Yuan at 2024/01/23 17:11
// leetgo: dev
// https://leetcode.cn/problems/delete-node-in-a-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func deleteNode(node *ListNode) {
	// 获取下一个节点的值
	next := node.Next
	// 把 node 转换为 next，然后删除 next
	node.Val = next.Val
	node.Next = next.Next
	next.Next = nil
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	node := Deserialize[int](ReadLine(stdin))
	deleteNode(head, node)
	ans := head

	fmt.Println("\noutput:", Serialize(ans))
}
