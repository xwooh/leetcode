// Created by Yuan at 2023/11/27 21:43
// leetgo: dev
// https://leetcode.cn/problems/merge-two-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	h := &ListNode{}
	p := h

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			h.Next = list1
			list1 = list1.Next
		} else {
			h.Next = list2
			list2 = list2.Next
		}
		h = h.Next
	}

	if list1 != nil {
		h.Next = list1
	} else if list2 != nil {
		h.Next = list2
	}

	ans = p.Next

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	list1 := Deserialize[*ListNode](ReadLine(stdin))
	list2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := mergeTwoLists(list1, list2)

	fmt.Println("\noutput:", Serialize(ans))
}
