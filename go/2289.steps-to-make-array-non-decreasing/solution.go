// Created by Yuan at 2023/10/30 22:17
// leetgo: dev
// https://leetcode.cn/problems/steps-to-make-array-non-decreasing/

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

type E struct {
	V    int
	Step int
}

func totalSteps(nums []int) (ans int) {
	stack := []E{}

	var c int
	for i := len(nums) - 1; i >= 0; i-- {
		c = 0
		for len(stack) > 0 && nums[i] > stack[len(stack)-1].V {
			top := stack[len(stack)-1]
			c += top.Step
			ans = max(ans, c)
			stack = stack[:len(stack)-1]
		}
		// 清零，重新计算
		if len(stack) == 0 {
			stack = append(stack, E{nums[i], 1})
		} else {
			stack = append(stack, E{nums[i], max(c, 1)})
		}
	}

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := totalSteps(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
