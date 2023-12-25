// Created by Yuan at 2023/12/25 10:50
// leetgo: dev
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	// 左子树为空，右子树也就绪了，则直接返回
	if root.Left == nil {
		return
	}

	// ln 的最后右节点上接上 rn
	rn := root.Right
	ln := root.Left
	for ln.Right != nil {
		ln = ln.Right
	}
	ln.Right = rn

	// root 右节点先从左节点开始
	root.Right = root.Left
	root.Left = nil
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	flatten(root)
	ans := root

	fmt.Println("\noutput:", Serialize(ans))
}
