// Created by Yuan at 2023/10/14 22:33
// leetgo: dev
// https://leetcode.cn/problems/search-in-rotated-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func search(nums []int, target int) (ans int) {
	// 先找断点 k
	l, r := 0, len(nums)-1

	for l < r && nums[l] > nums[r] {
		m := (l + r) / 2

		// 等号确保 l、m 重合时 l 右移
		if nums[l] <= nums[m] {
			// 在右侧
			l = m + 1
		} else {
			// 在左侧
			r = m
		}
	}

	// 再找 target
	var lx, rx int
	if l == 0 {
		// 递增序列
		lx, rx = 0, len(nums)-1
	} else {
		// 根据断点判断 target 在左侧还是右侧
		if nums[0] <= target && target <= nums[l-1] {
			// target 在左侧
			lx, rx = 0, l-1
		} else {
			lx, rx = l, len(nums)-1
		}
	}

	for lx < rx {
		mx := (lx + rx) / 2

		if target == nums[mx] {
			return mx
		} else if target > nums[mx] {
			// 右移
			lx = mx + 1
		} else {
			// 左移
			rx = mx - 1
		}
	}

	if nums[lx] == target {
		return lx
	}

	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := search(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
