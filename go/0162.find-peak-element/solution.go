// Created by Yuan at 2023/12/01 17:32
// leetgo: dev
// https://leetcode.cn/problems/find-peak-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func findPeakElement(nums []int) (ans int) {
	if len(nums) <= 1 {
		ans = len(nums) - 1
		return
	}

	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2 + 1
		if nums[m] > nums[m-1] {
			// 峰值在右侧
			l = m
		} else if nums[m] < nums[m-1] {
			// 峰值在左侧
			r = m - 1
		}
	}

	ans = l
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findPeakElement(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
