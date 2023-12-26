// Created by Yuan at 2023/12/25 18:26
// leetgo: dev
// https://leetcode.cn/problems/sort-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func findMid(head *ListNode) *ListNode {
	s, f := head, head.Next
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}

func merge(ln, rn *ListNode) *ListNode {
	h := &ListNode{}
	p := h

	for ln != nil && rn != nil {
		if ln.Val < rn.Val {
			h.Next = ln
			h = h.Next

			ln = ln.Next
		} else {
			h.Next = rn
			h = h.Next

			rn = rn.Next
		}
	}

	for ln != nil {
		h.Next = ln
		h = h.Next
		ln = ln.Next
	}

	for rn != nil {
		h.Next = rn
		h = h.Next
		rn = rn.Next
	}

	return p.Next
}

func sortList(head *ListNode) (ans *ListNode) {
	if head == nil {
		return
	} else if head.Next == nil {
		ans = head
		return
	}

	m := findMid(head)
	l := head
	r := m.Next
	m.Next = nil

	l = sortList(l)
	r = sortList(r)

	ans = merge(l, r)

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := sortList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
