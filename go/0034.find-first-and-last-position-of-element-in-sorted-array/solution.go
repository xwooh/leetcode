// Created by Yuan at 2023/10/15 11:54
// leetgo: dev
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchRange(nums []int, target int) (ans []int) {

	// 找两次，一次左边界，一次右边界
	ans = []int{-1, -1}

	if len(nums) == 0 {
		return
	}

	// 先找左边界
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2

		if nums[m] > target {
			// 左移
			r = m - 1
		} else if nums[m] < target {
			// 右移
			l = m + 1
		} else {
			// 左移
			// 因为找左边界，所以等于的时候不要移位了
			r = m
		}
	}
	if nums[l] == target {
		ans[0] = l
	}

	// 再找右边界
	l, r = 0, len(nums)-1
	for l < r {
		// 因为 等于 的时候没有右移，所以这边 m 要往右偏一格
		m := (l+r)/2 + 1

		if nums[m] > target {
			// 左移
			r = m - 1
		} else if nums[m] < target {
			// 右移
			// m 已经偏移过了，就不要偏移了
			l = m
		} else {
			// 右移
			// 因为找右边界，所以等于的时候不要移位了
			l = m
		}
	}
	if nums[l] == target {
		ans[1] = l
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchRange(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
