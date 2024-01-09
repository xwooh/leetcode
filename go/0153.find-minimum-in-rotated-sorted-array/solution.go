// Created by Yuan at 2023/12/18 16:59
// leetgo: dev
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
func fundS(nums []int) int {
	// 找断点处
	// 比如: 6, 7, 0, 1, 2 断点为 0
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2

		if nums[m] < nums[r] {
			// 说明 m 在断点右边 => 左移
			r = m
		} else if nums[m] > nums[r] {
			// m 在断点左侧 => 右移
			l = m + 1
		} else if nums[m] > nums[r] {
			return m
		}
	}

	return l
}

func findMin(nums []int) (ans int) {
	if len(nums) <= 0 {
		return
	} else if len(nums) == 1 {
		return nums[0]
	}

	a, b := nums[0], nums[len(nums)-1]
	if a < b {
		// 此时整个数组为升序
		ans = nums[0]
	} else {
		// 中间有断点，但是前后都是升序
		// 最小值即在断点处
		idx := fundS(nums)
		ans = nums[idx]
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findMin(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
