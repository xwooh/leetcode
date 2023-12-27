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
	h1, h2 := headA, headB

	// 如果有交点则在交点处退出
	// 如果没交点，因为 A+B 长度 == B+A 长度，所以最终也会在 nil 处相遇(相等)
	for h1 != h2 {
		if h1 == nil {
			// h1 走完了，就去到 headB 起始处
			h1 = headB
		} else {
			h1 = h1.Next
		}

		if h2 == nil {
			// h2 走完了，就去到 headA 起始处
			h2 = headA
		} else {
			h2 = h2.Next
		}

	}

	ans = h1

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
