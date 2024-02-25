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
func quickSearch(nums []int, target int) int {
	for idx, n := range nums {
		if n == target {
			return idx
		}
	}
	return -1
}

func findStart(nums []int) int {
	// 找到断点
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] <= nums[len(nums)-1] {
			// 左移
			r = m
		} else if nums[m] > nums[len(nums)-1] {
			// 右移
			l = m + 1
		}
	}

	if l == 0 {
		return l
	}

	if nums[l-1] > nums[l] {
		return l
	}

	return -1
}

func bs(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			// 左侧
			r = m - 1
		} else {
			// 右侧
			l = m + 1
		}
	}

	if nums[l] == target {
		return l
	}

	return -1
}

func search(nums []int, target int) (ans int) {
	if len(nums) <= 100 {
		return quickSearch(nums, target)
	}

	// 找到断点处
	s := findStart(nums)

	// 判断 target 在断点的左侧还是右侧
	if nums[s] == target {
		return s
	} else if nums[s] > target {
		// 最小值都比 target 大，那就不存在 target
		return -1
	} else if nums[len(nums)-1] < target {
		// 在左侧
		return bs(nums[:s], target)
	} else {
		// 在右侧
		r := bs(nums[s:], target)
		if r == -1 {
			return -1
		}
		return s + r
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := search(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
