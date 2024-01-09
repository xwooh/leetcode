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

	// 在右节点找公共祖先
	r := lowestCommonAncestor(root.Right, p, q)
	// 在左节点找公共祖先
	l := lowestCommonAncestor(root.Left, p, q)

	if r == nil {
		// 右节点里面没找到公共祖先
		//	那公共祖先就在左节点里面
		return l
	}

	if l == nil {
		return r
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
