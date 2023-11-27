// Created by Yuan at 2023/11/27 22:20
// leetgo: dev
// https://leetcode.cn/problems/merge-k-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func mergeTwo(a, b *ListNode) *ListNode {
	h := &ListNode{}
	p := h
	for a != nil && b != nil {
		if a.Val > b.Val {
			p.Next = b
			b = b.Next
		} else {
			p.Next = a
			a = a.Next
		}
		p = p.Next
	}
	if a != nil {
		p.Next = a
	} else if b != nil {
		p.Next = b
	}
	return h.Next
}

func mergeKLists(lists []*ListNode) (ans *ListNode) {
	if len(lists) == 0 {
		return
	} else if len(lists) == 1 {
		ans = lists[0]
		return
	} else if len(lists) == 2 {
		ans = mergeTwo(lists[0], lists[1])
		return
	}

	m := len(lists) / 2
	l := mergeKLists(lists[:m])
	r := mergeKLists(lists[m:])

	ans = mergeTwo(l, r)

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	lists := Deserialize[[]*ListNode](ReadLine(stdin))
	ans := mergeKLists(lists)

	fmt.Println("\noutput:", Serialize(ans))
}
