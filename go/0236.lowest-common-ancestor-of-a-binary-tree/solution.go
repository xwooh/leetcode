// Created by Yuan at 2023/11/01 17:56
// leetgo: dev
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	right := lowestCommonAncestor(root.Right, p, q)
	left := lowestCommonAncestor(root.Left, p, q)

	if right == nil {
		// 右子树没找到，却在左子树找到了，那说明找到的就是最近公共祖先
		return left
	} else if left == nil {
		return right
	}
	return root
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	p := Deserialize[int](ReadLine(stdin))
	q := Deserialize[int](ReadLine(stdin))
	ans := lowestCommonAncestor(root, p, q)

	fmt.Println("\noutput:", Serialize(ans))
}
