// Created by Yuan at 2023/11/27 22:35
// leetgo: dev
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeNthFromEnd(head *ListNode, n int) (ans *ListNode) {
	// 快慢指针，先让快指针先走 n 步
	s := &ListNode{Next: head}
	fn := s
	sn := s

	c := 0
	for fn.Next != nil && c < n {
		fn = fn.Next
		c++
	}

	for fn.Next != nil {
		fn = fn.Next
		sn = sn.Next
	}

	// 删除 sn.Next
	sn.Next = sn.Next.Next

	ans = s.Next
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := removeNthFromEnd(head, n)

	fmt.Println("\noutput:", Serialize(ans))
}
