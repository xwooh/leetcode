// Created by Yuan at 2023/12/27 11:49
// leetgo: dev
// https://leetcode.cn/problems/linked-list-cycle-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func findCircleNode(head *ListNode) *ListNode {
	s, f := head, head
	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return s
		}
	}

	return nil
}

func detectCycle(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 先用快慢指针判断是否有环
	c := findCircleNode(head)
	if c == nil {
		return
	}

	// 再来一个从 head 出发的指针，和相遇点一起往前走，相遇点就是环的起点
	ans = head
	for ans != c {
		ans = ans.Next
		c = c.Next
	}

	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	pos := Deserialize[int](ReadLine(stdin))
	ans := detectCycle(head, pos)

	fmt.Println("\noutput:", Serialize(ans))
}
