// Created by Yuan at 2023/12/27 16:41
// leetgo: dev
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getIntersectionNode(headA, headB *ListNode) (ans *ListNode) {
	if headA == nil || headB == nil {
		return
	}

	h1, h2 := headA, headB

	for h1 != nil && h2 != nil {
		if h1 == h2 {
			return h1
		}

		h1 = h1.Next
		h2 = h2.Next
	}

	if h1 == nil {
		h1 = headB
	}
	if h2 == nil {
		h2 = headA
	}

	for h1 != nil && h2 != nil {
		if h1 == h2 {
			return h1
		}

		h1 = h1.Next
		h2 = h2.Next
	}

	if h1 == nil {
		h1 = headB
	}
	if h2 == nil {
		h2 = headA
	}

	for h1 != nil && h2 != nil {
		if h1 == h2 {
			return h1
		}

		h1 = h1.Next
		h2 = h2.Next
	}

	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	intersectVal := Deserialize[int](ReadLine(stdin))
	listA := Deserialize[*ListNode](ReadLine(stdin))
	listB := Deserialize[*ListNode](ReadLine(stdin))
	skipA := Deserialize[int](ReadLine(stdin))
	skipB := Deserialize[int](ReadLine(stdin))
	ans := getIntersectionNode(intersectVal, listA, listB, skipA, skipB)

	fmt.Println("\noutput:", Serialize(ans))
}
