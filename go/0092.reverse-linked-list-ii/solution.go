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
	p := 1
	dummy := &ListNode{Next: head}

	var x *ListNode
	if p == left {
		x = dummy
	}

	for p < right && head != nil {
		if p+1 == left {
			x = head
		}

		if p < left {
			head = head.Next
		} else {
			// 反转链表套路
			tmp := head.Next
			head.Next = head.Next.Next
			tmp.Next = x.Next
			x.Next = tmp
		}

		// 前进
		p++
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
