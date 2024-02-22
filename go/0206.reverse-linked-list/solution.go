// Created by Yuan at 2023/11/27 13:34
// leetgo: dev
// https://leetcode.cn/problems/reverse-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func reverseList(head *ListNode) (ans *ListNode) {
	if head == nil || head.Next == nil {
		ans = head
		return
	}

	dummy := &ListNode{Next: head}
	p := head

	for p.Next != nil {
		// 取后一个数
		tmp := p.Next

		// p 指向 tmp 后面的节点
		p.Next = p.Next.Next

		// tm 插入到头结点前面
		tmp.Next = dummy.Next
		dummy.Next = tmp
	}

	ans = dummy.Next
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
