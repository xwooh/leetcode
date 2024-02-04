// Created by Yuan at 2023/11/27 13:34
// leetgo: dev
// https://leetcode.cn/problems/palindrome-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func getRightHalfNode(head *ListNode) *ListNode {
	s, f := head, head
	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	r := s.Next

	// 反转链表 r
	dummy := &ListNode{Next: r}
	p := r
	for p.Next != nil {
		tmp := p.Next
		p.Next = p.Next.Next
		tmp.Next = dummy.Next
		dummy.Next = tmp
	}

	return dummy.Next
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 判断是否是回文链表
	// 快慢指针找到链表的中点
	// - 奇数链表，当快指针 .Next == nil 时，慢指针正好走到中点，即 .Next = 另一半链表的头
	//   1 -> 3 -> 5 -> 3 -> 1
	// - 偶数链表，当快指针 .Next.Next =nil 时，慢指针正好走到一半的末尾，即 .Next = 另一半链表的头
	//   1 -> 3 -> 5 -> 5 -> 3 -> 1
	r := getRightHalfNode(head)
	for r != nil {
		if r.Val != head.Val {
			return false
		}

		r = r.Next
		head = head.Next
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := isPalindrome(head)

	fmt.Println("\noutput:", Serialize(ans))
}
