// Created by Yuan at 2023/11/06 17:38
// leetgo: dev
// https://leetcode.cn/problems/search-a-2d-matrix-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchMatrix(matrix [][]int, target int) bool {

	rl := len(matrix)
	cl := len(matrix[0])

	var r int
	for r >= 0 && r < rl {
		// 先往右，确定右侧范围
		// 再不断往下找
		c := 0
		for c >= 0 && c < cl {
			v := matrix[r][c]
			if v == target {
				return true
			} else if v < target {
				c++
			} else if v > target {
				break
			}
		}

		cl = c
		r++

	}

	return false

}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchMatrix(matrix, target)

	fmt.Println("\noutput:", Serialize(ans))
}
