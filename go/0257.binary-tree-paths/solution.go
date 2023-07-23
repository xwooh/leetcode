// Created by Yuan at 2023/07/23 14:57
// leetgo: dev
// https://leetcode.cn/problems/binary-tree-paths/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func binaryTreePaths(root *TreeNode) (ans []string) {
	if root == nil {
		return ans
	}

	if root.Left == nil && root.Right == nil {
		ans := append(ans, strconv.Itoa(root.Val))
		return ans
	}

	lAns := binaryTreePaths(root.Left)
	rAns := binaryTreePaths(root.Right)

	// 左边的所有结果 + 右边所有结果
	for _, s := range lAns {
		ns := strconv.Itoa(root.Val) + "->" + s
		ans = append(ans, ns)
	}
	for _, s := range rAns {
		ns := strconv.Itoa(root.Val) + "->" + s
		ans = append(ans, ns)
	}

	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := binaryTreePaths(root)

	fmt.Println("\noutput:", Serialize(ans))
}
