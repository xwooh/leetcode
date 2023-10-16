// Created by Yuan at 2023/10/16 23:00
// leetgo: dev
// https://leetcode.cn/problems/diameter-of-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TreeMaxDepth(root *TreeNode, maxDiameter *int) int {
	// maxDiameter 当前树的最大直径
	if root == nil {
		return 0
	}

	l := TreeMaxDepth(root.Left, maxDiameter)
	r := TreeMaxDepth(root.Right, maxDiameter)

	// 当前的最大直径 = 左节点最大深度 + 右节点最大深度
	currentDiameter := l + r
	if currentDiameter > *maxDiameter {
		*maxDiameter = currentDiameter
	}

	return max(l, r) + 1
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}

	var maxD int
	TreeMaxDepth(root, &maxD)

	return maxD
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := diameterOfBinaryTree(root)

	fmt.Println("\noutput:", Serialize(ans))
}
