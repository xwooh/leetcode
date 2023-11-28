// Created by Yuan at 2023/11/27 14:58
// leetgo: dev
// https://leetcode.cn/problems/partition-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

func sortByTwoListNode(head *ListNode, x int) (ans *ListNode) {
	// < x 的追加到 sn 后面，>= x 的追加到 bn 后面
	sn := &ListNode{}
	bn := &ListNode{}

	sp := sn
	bp := bn

	for head != nil {
		if head.Val < x {
			sp.Next = &ListNode{Val: head.Val, Next: nil}
			sp = sp.Next
		} else {
			bp.Next = &ListNode{Val: head.Val, Next: nil}
			bp = bp.Next
		}
		head = head.Next
	}

	ans = sn.Next
	if ans == nil {
		ans = bn.Next
	} else {
		ap := ans
		for ap.Next != nil {
			ap = ap.Next
		}
		ap.Next = bn.Next
	}

	return
}

// @lc code=begin
func partition(head *ListNode, x int) (ans *ListNode) {
	ans = sortByTwoListNode(head, x)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	x := Deserialize[int](ReadLine(stdin))
	ans := partition(head, x)

	fmt.Println("\noutput:", Serialize(ans))
}
