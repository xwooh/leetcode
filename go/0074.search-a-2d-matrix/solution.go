// Created by Yuan at 2023/10/14 16:37
// leetgo: dev
// https://leetcode.cn/problems/search-a-2d-matrix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	// 先二分搜索第一列
	l, r := 0, len(matrix)-1
	// 循环出去的时候 l == r，没有越界
	for l < r {
		row := (l + r) / 2
		rv := matrix[row][0]
		if rv == target {
			return true
		} else if rv < target {
			// 下移
			l = row + 1 // row 肯定不是解，所以 +1
		} else {
			// 上移
			r = row - 1 // row 肯定不是解，所以 -1
		}
	}

	var nums []int
	if matrix[l][0] > target {
		// 在上一行
		if l == 0 {
			return false
		}
		nums = matrix[l-1]
	} else {
		// 在这行
		nums = matrix[l]
	}

	// 再二分搜索对应行
	l, r = 0, len(nums)-1
	// 循环出去的时候 l > r，此时已经越界，所以循环出去肯定没解
	for l <= r {
		col := (l + r) / 2
		cv := nums[col]
		if cv == target {
			return true
		} else if cv < target {
			// 右移
			l = col + 1
		} else {
			// 左移
			r = col - 1
		}
	}

	return false
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchMatrix(matrix, target)

	fmt.Println("\noutput:", Serialize(ans))
}
