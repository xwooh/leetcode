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
func reverseNode(head *ListNode) *ListNode {
	p := &ListNode{Next: head}

	for head != nil && head.Next != nil {
		tmp := head.Next
		head.Next = head.Next.Next
		tmp.Next = p.Next
		p.Next = tmp
	}

	return p.Next
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 取后半段节点
	s, f := head, head
	for f != nil && f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	n := s.Next

	// 后半段节点反转
	n = reverseNode(n)

	// 判断是否回文
	h1, h2 := head, n
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1 = h1.Next
		h2 = h2.Next
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
