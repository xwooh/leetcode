// Created by Yuan at 2023/12/27 11:06
// leetgo: dev
// https://leetcode.cn/problems/linked-list-cycle/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hasCycle(head *ListNode) bool {
	// 如果无环，肯定有终点
	// 如果有环，快慢指针肯定会在某个节点相遇

	s, f := head, head

	for s != nil && f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			return true
		}
	}
	return false
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	pos := Deserialize[int](ReadLine(stdin))
	ans := hasCycle(head, pos)

	fmt.Println("\noutput:", Serialize(ans))
}
