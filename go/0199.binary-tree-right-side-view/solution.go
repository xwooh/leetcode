// Created by Yuan at 2024/01/09 16:18
// leetgo: dev
// https://leetcode.cn/problems/binary-tree-right-side-view/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		// 最后一个就是最右边
		top := q[len(q)-1]
		ans = append(ans, top.Val)

		// 加入下一行节点
		nextQ := []*TreeNode{}
		for len(q) > 0 {
			head := q[0]
			q = q[1:]
			if head.Left != nil {
				nextQ = append(nextQ, head.Left)
			}
			if head.Right != nil {
				nextQ = append(nextQ, head.Right)
			}
		}
		q = nextQ
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := rightSideView(root)

	fmt.Println("\noutput:", Serialize(ans))
}
