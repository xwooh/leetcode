// Created by Yuan at 2024/01/08 14:35
// leetgo: dev
// https://leetcode.cn/problems/reverse-nodes-in-k-group/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseKGroup(head *ListNode, k int) (ans *ListNode) {
	if k == 1 {
		ans = head
		return
	}

	// 先计算下能跑几遍
	tc := 0
	for x := head; x != nil; x = x.Next {
		tc++
	}
	ts := tc / k
	if ts <= 0 {
		ans = head
		return
	}

	// 做 ts 遍
	dummy := &ListNode{Next: head}

	// 哨兵节点
	h := dummy
	// 当前节点
	p := dummy

	c := 0
	rc := 0
	for p.Next != nil {
		if rc == 0 {
			p = p.Next
		} else {
			// 需要把这个节点插入到头结点前面
			tmp := p.Next

			p.Next = p.Next.Next
			tmp.Next = h.Next
			h.Next = tmp
		}

		// 做完一个节点，计数 +1
		rc++

		if rc == k {
			// 重置 rc 计数
			rc = 0
			h = p

			// 完整地做了一遍
			c++
		}

		if c == ts {
			// 后面的已经不够做一遍了，终止
			break
		}

	}

	ans = dummy.Next

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := reverseKGroup(head, k)

	fmt.Println("\noutput:", Serialize(ans))
}
