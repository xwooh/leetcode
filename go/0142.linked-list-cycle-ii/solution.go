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

func detectCycle(head *ListNode) (ans *ListNode) {
	// 1. 找出相遇点
	// 2. 此时再让快指针回到起点，速度和慢指针一致
	// 3. 再次相遇即为相交起点

	s, f := head, head

	for s != nil && f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			break
		}
	}

	if s == nil || f == nil || f.Next == nil {
		return
	}

	f = head
	for s != f {
		s = s.Next
		f = f.Next
	}

	ans = s

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
