// Created by Yuan at 2024/01/08 13:11
// leetgo: dev
// https://leetcode.cn/problems/reverse-linked-list-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseBetween(head *ListNode, left int, right int) (ans *ListNode) {
	// 就一个节点，没法翻转
	if left == right {
		ans = head
		return
	}

	dummy := &ListNode{Next: head}
	// 当前节点
	h := dummy
	// 「头」结点
	p := dummy
	c := 0
	for c <= right && h.Next != nil {
		if c+1 < left {
			p = p.Next
		}

		if c < left {
			h = h.Next
		} else if c > left {
			// 后面的节点都要移到 h 前面
			tmp := h.Next

			h.Next = h.Next.Next
			tmp.Next = p.Next
			p.Next = tmp
		}
		c++
	}

	ans = dummy.Next

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	left := Deserialize[int](ReadLine(stdin))
	right := Deserialize[int](ReadLine(stdin))
	ans := reverseBetween(head, left, right)

	fmt.Println("\noutput:", Serialize(ans))
}
