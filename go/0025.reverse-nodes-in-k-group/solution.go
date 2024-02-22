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
	dynamicDummy := &ListNode{Next: head}
	dummy := dynamicDummy

	// dynamicDummy 指向上一次的尾结点
	// 先遍历一次，看看后面的节点数是否满足 k 个
	// 如果满足 k 个再去翻转

	p := head
	for p != nil && p.Next != nil {
		t := p
		c := 0
		for t != nil && c < k {
			t = t.Next
			c++
		}

		if c < k {
			// 后面节点不足，直接退出
			break
		}

		// 后面节点足够，那就翻转
		// 翻转 k - 1 次
		for c > 1 {
			tmp := p.Next
			p.Next = p.Next.Next
			tmp.Next = dynamicDummy.Next
			dynamicDummy.Next = tmp
			c--
		}

		// 翻转完成后 p 指向的是此次 k 个节点的尾结点
		dynamicDummy = p
		p = p.Next
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
