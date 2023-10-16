// Created by Yuan at 2023/10/16 22:22
// leetgo: dev
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func DivDepth(root *TreeNode) int {
	// 单个节点的情况
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}

	// 左右节点的各自情况
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)

	// 整个高度 = max(Left, Right) + 1 // (根节点)
	if ld > rd {
		return ld + 1
	} else {
		return rd + 1
	}
}

func maxDepth(root *TreeNode) int {
	return DivDepth(root)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := maxDepth(root)

	fmt.Println("\noutput:", Serialize(ans))
}
