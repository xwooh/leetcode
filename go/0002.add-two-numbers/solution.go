// Created by Yuan at 2024/01/27 16:00
// leetgo: dev
// https://leetcode.cn/problems/add-two-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func add(sum int, step *int, node *ListNode) *ListNode {
	// 本次计算后的个位数
	v := sum % 10
	// 本次计算后的进位数
	*step = sum / 10

	x := &ListNode{Val: v}
	node.Next = x
	return node.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	r := &ListNode{}
	head := r

	var step int
	for l1 != nil && l2 != nil {
		r = add(l1.Val+l2.Val+step, &step, r)
		l1 = l1.Next
		l2 = l2.Next
	}

	// l2 已经每值了，l1 还有值
	for l1 != nil {
		r = add(l1.Val+step, &step, r)
		l1 = l1.Next
	}

	// l1 已经每值了，l2 还有值
	for l2 != nil {
		r = add(l2.Val+step, &step, r)
		l2 = l2.Next
	}

	// 最后还有进位
	if step != 0 {
		s := 0
		r = add(step, &s, r)
	}

	return head.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	l1 := Deserialize[*ListNode](ReadLine(stdin))
	l2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := addTwoNumbers(l1, l2)

	fmt.Println("\noutput:", Serialize(ans))
}
