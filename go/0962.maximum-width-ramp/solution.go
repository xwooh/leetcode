// Created by Yuan at 2024/01/26 15:26
// leetgo: dev
// https://leetcode.cn/problems/maximum-width-ramp/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type P struct {
	v   int
	idx int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxWidthRamp(nums []int) (ans int) {
	// 严格单调递减栈
	stack := []P{}
	for idx, v := range nums {
		for len(stack) == 0 || len(stack) > 0 && v < stack[len(stack)-1].v {
			stack = append(stack, P{v, idx})
		}
	}

	// 从右往左，查找最宽坡度
	for r := len(nums) - 1; r >= 0; r-- {
		for len(stack) > 0 && stack[len(stack)-1].v <= nums[r] {
			top := stack[len(stack)-1]
			ans = max(ans, r-top.idx)
			stack = stack[:len(stack)-1]
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maxWidthRamp(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
